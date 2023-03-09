package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

var (
	urlTlaloc   string
	tokenTlaloc string
)

func NewTlaloc(configuration *models.Configuration, requests *Requests) *Tlaloc {
	cs := Tlaloc{Configuration: configuration, Requests: requests}
	cs.init()
	return &cs
}

type Tlaloc struct {
	Configuration *models.Configuration
	Requests      *Requests
}

func (tlaloc Tlaloc) init() {
	// url al recurso de tlaloc
	urlTlaloc = fmt.Sprintf("%s/%s", strings.TrimSuffix(tlaloc.Configuration.APITlaloc, "/"), constants.TlalocLocation)
	tokenTlaloc = tlaloc.Configuration.TokenTlaloc
}

func (tlaloc Tlaloc) GetLocTlaloc() (map[string]models.TlalocLocation, error) {
	headers := make(http.Header)
	headers.Add("Authorization", tokenTlaloc)
	response, body, err := tlaloc.Requests.SendRequestAndBody(interfaces.RequestData{
		URL:     urlTlaloc,
		Method:  "GET",
		Headers: headers,
	})
	if err != nil {
		return make(map[string]models.TlalocLocation), err
	}
	response.Body.Close()
	if response.StatusCode != http.StatusOK {
		logger.LogSimpleDebug(tokenTlaloc)
		return make(map[string]models.TlalocLocation), errors.New(tlaloc.getCodeStatusError(response))
	}
	tlalocResponse := models.TlalocResponse{}
	err = json.Unmarshal(body, &tlalocResponse)
	if err != nil {
		return make(map[string]models.TlalocLocation), err
	}
	locations := make(map[string]models.TlalocLocation)
	for _, element := range tlalocResponse.Locations {
		locations[element.Compatibility.Locations] = element
	}
	return locations, err
}

func (tlaloc Tlaloc) getCodeStatusError(response *http.Response) string {
	return "tlaloc error, status code invalid. URL: " + urlTlaloc + ", code: " + strconv.Itoa(response.StatusCode)
}

func TlalocCheckHealth(httpClient *http.Client, apiTlaloc string) error {
	urlTlalocHealth := fmt.Sprintf("%s/tlaloc/health", strings.TrimSuffix(apiTlaloc, "/"))
	chanTlalocHealth := make(chan string)
	defer closeChannels(chanTlalocHealth)
	go processHealth(urlTlalocHealth, httpClient, chanTlalocHealth)
	messageHealth := <-chanTlalocHealth
	return concatErrors(messageHealth)
}
