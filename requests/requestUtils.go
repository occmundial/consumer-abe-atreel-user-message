package requests

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// processHealth : regresa el statusCode del endpoint de salud
func processHealth(urlAPI string, httpClient *http.Client, chanErrorMsg chan string) {
	ctx := context.Background()
	req2, _ := http.NewRequestWithContext(ctx, http.MethodGet, urlAPI, http.NoBody)
	response, err := httpClient.Do(req2)
	response.Body.Close()
	if err != nil {
		chanErrorMsg <- fmt.Sprintf("Service Unavailable -> %s", err.Error())
	} else if response.StatusCode != http.StatusOK {
		chanErrorMsg <- fmt.Sprintf("Service Unavailable -> %s", urlAPI)
	} else {
		chanErrorMsg <- ""
	}
}

func concatErrors(apiErrors ...string) error {
	errorMsgs := ""
	for _, msg := range apiErrors {
		if len(msg) > 0 {
			errorMsgs += msg + "\n"
		}
	}
	if len(errorMsgs) > 0 {
		return errors.New(strings.TrimSuffix(errorMsgs, "\n"))
	}
	return nil
}

func closeChannels(channels ...chan string) {
	for _, item := range channels {
		close(item)
	}
}
