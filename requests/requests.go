package requests

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

func NewRequests(configuration *models.Configuration) *Requests {
	return &Requests{HTTPClient: libs.InitHTTPClient(configuration)}
}

type Requests struct {
	HTTPClient *http.Client
}

func (requests Requests) SendRequestAndBody(r interfaces.RequestData) (*http.Response, []byte, error) {
	headers := r.Headers
	if headers == nil {
		headers = make(http.Header)
	}
	parsedURL, err := url.Parse(r.URL)
	if err != nil {
		return nil, []byte{}, err
	}
	request := new(http.Request)
	request.Header = headers
	if r.ContentType != "" {
		request.Header.Add("Content-Type", r.ContentType)
	}
	request.Method = r.Method
	request.URL = parsedURL
	if r.Data != "" {
		stringReader := strings.NewReader(r.Data)
		request.Body = io.NopCloser(stringReader)
	}
	response, err := new(http.Client).Do(request)
	if err != nil {
		return response, []byte{}, err
	}
	body, err := getBody(response)
	return response, body, err
}

func getBody(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func (requests Requests) CheckHealth(urlRequest string) error {
	chanRequestHealth := make(chan string)
	defer closeChannels(chanRequestHealth)
	go processHealth(urlRequest, requests.HTTPClient, chanRequestHealth)
	messageHealth := <-chanRequestHealth
	return concatErrors(messageHealth)
}
