package libs

import (
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

func InitRetryHTTPClient(config *models.Configuration) *http.Client {
	client := retryablehttp.NewClient()
	client.Logger = nil
	client.HTTPClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	client.RetryWaitMin = time.Millisecond * time.Duration(config.RetryWaitMin)
	client.RetryWaitMax = time.Millisecond * time.Duration(config.RetryWaitMax)
	client.RetryMax = config.RetryMax
	client.HTTPClient.Timeout = time.Second * time.Duration(config.APITimeout)
	return client.StandardClient()
}

// InitHTTPClient :
func InitHTTPClient(config *models.Configuration) *http.Client {
	return &http.Client{
		Timeout: time.Second * time.Duration(config.APITimeout),
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
}
