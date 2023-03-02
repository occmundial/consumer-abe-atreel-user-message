package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/database"
	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

var (
	urlAtreel       string
	urlAtreelHealth string
	stateDic        map[string]string
)

func NewAtreel(configuration *models.Configuration, queries *database.Queries) *Atreel {
	retryHTTPClient := libs.InitRetryHTTPClient(configuration)
	httpClient := libs.InitHTTPClient(configuration)
	cs := Atreel{Configuration: configuration, RetryHTTPClient: retryHTTPClient, HTTPClient: httpClient}
	cs.init(queries)
	return &cs
}

type Atreel struct {
	Configuration   *models.Configuration
	RetryHTTPClient *http.Client
	HTTPClient      *http.Client
}

func (atreel Atreel) init(queries interfaces.IQuery) {
	urlAtreel = fmt.Sprintf("%s/atreel/v3/emails", strings.TrimSuffix(atreel.Configuration.APIAtreel, "/"))
	urlAtreelHealth = fmt.Sprintf("%s/health", strings.TrimSuffix(atreel.Configuration.APIAtreel, "/"))
	var err error
	stateDic, err = queries.GetDicState()
	if err != nil {
		logger.Fatal("processAtreel", "init", err)
	}
}

func (atreel Atreel) PostCorreo(messageFromKafka *models.MessageToProcess) error {
	data := ConvertJSONToHTMLAbeData{messageFromKafka.Recommendations, messageFromKafka.Name, stateDic}
	jobsIds, dynamicTemplateData := ConvertJSONToHTMLABE(&data, atreel.Configuration)
	sendgridJSON := models.SendgridJSON{
		TemplateID: constants.TemplateID,
		JobID:      jobsIds,
		LoginID:    messageFromKafka.LoginID,
		Platform:   constants.Platform,
		Personalizations: []models.Personalizations{{
			To:                  []string{messageFromKafka.Email},
			DynamicTemplateData: dynamicTemplateData,
		}},
	}
	jsonBytes, e := json.Marshal(sendgridJSON)
	if e != nil {
		return e
	}
	ctx := context.Background()
	req2, _ := http.NewRequestWithContext(ctx, http.MethodPost, urlAtreel, bytes.NewBuffer(jsonBytes))
	req2.Header.Set("Content-Type", constants.JSONContentType)
	response, err := atreel.RetryHTTPClient.Do(req2)

	if err != nil {
		return err
	}
	response.Body.Close()
	if response.StatusCode != http.StatusCreated {
		return errors.New("invalid status code")
	}
	return nil
}

func AtreelCheckHealth(httpClient *http.Client) error {
	chanAtreelHealth := make(chan string)
	defer closeChannels(chanAtreelHealth)
	go processHealth(urlAtreelHealth, httpClient, chanAtreelHealth)
	messageHealth := <-chanAtreelHealth
	return concatErrors(messageHealth)
}
