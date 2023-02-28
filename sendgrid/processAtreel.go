package sendgrid

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/database"
	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/occmundial/consumer-abe-atreel-user-message/requests"
)

// NewAtreelProcessor
func NewAtreelProcessor(configuration *models.Configuration, queries *database.Queries, atreel *requests.Atreel) *AtreelProcessor {
	cs := AtreelProcessor{Configuration: configuration, Atreel: atreel}
	cs.init(queries)
	return &cs
}

// AtreelProcessor :
type AtreelProcessor struct {
	Configuration *models.Configuration
	Atreel        interfaces.IAtreel
}

func (processor *AtreelProcessor) init(queries interfaces.IQuery) {
	var err error
	if err != nil {
		logger.Fatal("processAtreel", "init", err)
	}
}

// isValidMessage: es un mensaje válido para la lógica de negocio
func IsValidMessage(message models.MessageToProcess) bool {
	return len(message.Email) > 0
}

// isValidEvent: es un mensaje válido del framework de eventos
func (processor *AtreelProcessor) CreateAndSendEmail(messageFromKafka *models.MessageToProcess) error {
	return processor.sendMessage(messageFromKafka)
}

func (processor *AtreelProcessor) sendMessage(messageFromKafka *models.MessageToProcess) error {
	return processor.Atreel.PostCorreo(messageFromKafka)
}
