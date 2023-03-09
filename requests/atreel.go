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
	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

var (
	urlAtreel       string
	urlAtreelHealth string
	locTlaloc       map[string]models.TlalocLocation
)

func NewAtreel(configuration *models.Configuration, tlaloc *Tlaloc) *Atreel {
	retryHTTPClient := libs.InitRetryHTTPClient(configuration)
	httpClient := libs.InitHTTPClient(configuration)
	cs := Atreel{Configuration: configuration, RetryHTTPClient: retryHTTPClient, HTTPClient: httpClient}
	cs.init(tlaloc)
	return &cs
}

type Atreel struct {
	Configuration   *models.Configuration
	RetryHTTPClient *http.Client
	HTTPClient      *http.Client
}

func (atreel Atreel) init(tlaloc interfaces.ITlaloc) {
	urlAtreel = fmt.Sprintf("%s/atreel/v3/emails", strings.TrimSuffix(atreel.Configuration.APIAtreel, "/"))
	urlAtreelHealth = fmt.Sprintf("%s/health", strings.TrimSuffix(atreel.Configuration.APIAtreel, "/"))
	var err error
	locTlaloc, err = tlaloc.GetLocTlaloc()
	if err != nil {
		logger.Fatal("requests", "init", err)
	}
}

func (atreel Atreel) PostCorreo(messageFromKafka *models.MessageToProcess) error {
	side := messageFromKafka.AbSide
	abTestName := messageFromKafka.AbTestName
	name := messageFromKafka.Name
	data := ConvertJSONToHTMLAbeData{messageFromKafka.Recommendations, name, locTlaloc, side, abTestName}
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
