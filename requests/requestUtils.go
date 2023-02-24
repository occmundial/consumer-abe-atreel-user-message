package requests

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"
)

// processHealth : regresa el statusCode del endpoint de salud
func processHealth(urlAPI string, chanErrorMsg chan string) {
	response, err := httpClient.Get(urlAPI)
	if err != nil {
		chanErrorMsg <- fmt.Sprintf("Service Unavailable -> %s", err.Error())
	} else if response.StatusCode != http.StatusOK {
		chanErrorMsg <- fmt.Sprintf("Service Unavailable -> %s", urlAPI)
	} else {
		chanErrorMsg <- ""
	}
}

func pingHealth(urlAPI string, chanErrorMsg chan string, timeout int) {
	host := strings.Split(urlAPI, "//")
	_, err := net.DialTimeout("tcp", host[1], time.Duration(timeout)*time.Second)
	if err == nil {
		chanErrorMsg <- ""
	} else {
		chanErrorMsg <- fmt.Sprintf("Service Unavailable -> %s", err.Error())
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
