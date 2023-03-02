package interfaces

import "net/http"

type IRequests interface {
	SendRequestAndBody(r RequestData) (*http.Response, []byte, error)
}

type RequestData struct {
	URL         string
	Headers     map[string][]string
	Method      string
	ContentType string
	Data        string
}
