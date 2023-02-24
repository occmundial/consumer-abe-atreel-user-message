package requests

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
)

func NewRequests() *Requests {
	return &Requests{}
}

type Requests struct {
}

func (requests Requests) SendRequestAndBody(r interfaces.RequestData) (*http.Response, []byte, error) {
	headers := r.Headers
	if headers == nil {
		headers = make(http.Header)
	}
	url, err := url.Parse(r.Url)
	if err != nil {
		return nil, []byte{}, err
	}
	request := new(http.Request)
	request.Header = headers
	request.Header.Add("Content-Type", r.ContentType)
	request.Method = r.Method
	request.URL = url
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
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func (request Requests) CheckHealth(urlRequest string) error {
	chanRequestHealth := make(chan string)
	defer closeChannels(chanRequestHealth)
	go processHealth(urlRequest, chanRequestHealth)
	messageHealth := <-chanRequestHealth
	return concatErrors(messageHealth)
}
