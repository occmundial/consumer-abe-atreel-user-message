package models

import (
	"github.com/segmentio/kafka-go"
)

// MessageForRead :
type MessageForRead struct {
	Key           string           `json:"key" example:"c4bf18e8-10f6-48ca-aa6f-61c7e8c25578"`
	Topic         string           `json:"topic"`
	Message       MessageToProcess `json:"message"`
	FlatMessage   string           `json:"flatMessage"`
	SourceMessage kafka.Message    `json:"-"`
}
