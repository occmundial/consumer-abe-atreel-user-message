package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/occmundial/consumer-abe-atreel-user-message/requests"

	"github.com/segmentio/kafka-go"
)

// getKafkaReader :
func getKafkaReader(config *models.Configuration) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        config.KafkaBrokers,
		GroupID:        config.GroupID,
		GroupTopics:    config.Topics,
		MinBytes:       10e3,        // 10KB
		MaxBytes:       10e6,        // 10MB
		CommitInterval: time.Second, // flushes commits to Kafka every second
	})
}

func kafkaMessageToMessageForRead(kafkaMessage kafka.Message) (models.MessageForRead, error) {
	message, err := deserializeMessage(kafkaMessage.Value)
	if err != nil {
		return models.MessageForRead{}, err
	}
	return models.MessageForRead{
		Key:           string(kafkaMessage.Key),
		Topic:         kafkaMessage.Topic,
		Message:       message,
		FlatMessage:   string(kafkaMessage.Value),
		SourceMessage: kafkaMessage,
	}, nil
}

func deserializeMessage(flatMessage []byte) (models.MessageToProcess, error) {
	message := models.MessageToProcess{}
	err := json.Unmarshal(flatMessage, &message)
	return message, err
}

// processHealth : regresa el statusCode del endpoint de salud
func processHealth(chanErrorMsg chan string, config *models.Configuration, queries interfaces.IQuery) {
	atreelError := requests.AtreelCheckHealth(config)
	dbError := queries.CheckHealth()
	chanErrorMsg <- getError(atreelError, dbError)
}

func getError(errs ...error) string {
	errorText := ""
	for _, err := range errs {
		if err != nil {
			if errorText != "" {
				errorText += ","
			}
			errorText += fmt.Sprintf("Service Unavailable -> %s,", err.Error())
		}
	}
	return errorText
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
