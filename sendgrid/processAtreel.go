package sendgrid

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/occmundial/consumer-abe-atreel-user-message/requests"
)

// NewAtreelProcessor
func NewAtreelProcessor(configuration *models.Configuration, atreel *requests.Atreel) *AtreelProcessor {
	cs := AtreelProcessor{Configuration: configuration, Atreel: atreel}
	return &cs
}

// AtreelProcessor :
type AtreelProcessor struct {
	Configuration *models.Configuration
	Atreel        interfaces.IAtreel
}

// isValidMessage: es un mensaje válido para la lógica de negocio
func IsValidMessage(message *models.MessageToProcess) bool {
	return len(message.Email) > 0
}

// isValidEvent: es un mensaje válido del framework de eventos
func (processor *AtreelProcessor) CreateAndSendEmail(messageFromKafka *models.MessageToProcess) error {
	return processor.sendMessage(messageFromKafka)
}

func (processor *AtreelProcessor) sendMessage(messageFromKafka *models.MessageToProcess) error {
	return processor.Atreel.PostCorreo(messageFromKafka)
}
