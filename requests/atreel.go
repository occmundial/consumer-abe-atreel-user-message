package requests

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/occmundial/consumer-abe-atreel-user-message/database"
	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"
	"net/http"
	"strings"

	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

var (
	urlAtreel       string
	urlAtreelHealth string
	stateDic        map[string]string
	retryHttpClient *http.Client
	httpClient      *http.Client
)

func NewAtreel(configuration *models.Configuration, queries *database.Queries) *Atreel {
	cs := Atreel{Configuration: configuration}
	cs.init(queries)
	return &cs
}

type Atreel struct {
	Configuration *models.Configuration
}

func (atreel Atreel) init(queries interfaces.IQuery) {
	retryHttpClient = libs.InitRetryHttpClient(atreel.Configuration)
	httpClient = libs.InitHttpClient(atreel.Configuration)
	urlAtreel = fmt.Sprintf("%s/atreel/v3/emails", strings.TrimSuffix(atreel.Configuration.APIAtreel, "/"))
	urlAtreelHealth = fmt.Sprintf("%s/health", strings.TrimSuffix(atreel.Configuration.APIAtreel, "/"))
	var err error
	stateDic, err = queries.GetDicState()
	if err != nil {
		logger.Fatal("processAtreel", "init", err)
	}
}

func (atreel Atreel) PostCorreo(messageFromKafka models.MessageToProcess) error {
	jobsIds, dynamicTemplateData := ConvertJsonToHtml_ABE(messageFromKafka.Recommendations, messageFromKafka.Name, stateDic, atreel.Configuration)
	sendgridJson := models.SendgridJson{
		Template_ID: constants.Template_ID,
		JobID:       jobsIds,
		LoginID:     messageFromKafka.LoginID,
		Platform:    constants.Platform,
		Personalizations: []models.Personalizations{{
			To:                  []string{messageFromKafka.Email},
			DynamicTemplateData: dynamicTemplateData,
		}},
	}
	jsonBytes, e := json.Marshal(sendgridJson)
	if e != nil {
		return e
	}
	response, err := retryHttpClient.Post(urlAtreel, constants.JSON_CONTENT_TYPE, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusCreated {
		return errors.New("invalid status code")
	}
	response.Body.Close()
	return nil
}

func AtreelCheckHealth(config *models.Configuration) error {
	chanAtreelHealth := make(chan string)
	defer closeChannels(chanAtreelHealth)
	go processHealth(urlAtreelHealth, chanAtreelHealth)
	messageHealth := <-chanAtreelHealth
	return concatErrors(messageHealth)
}
