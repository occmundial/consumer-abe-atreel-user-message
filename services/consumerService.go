package services

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/occmundial/consumer-abe-atreel-user-message/repositories"
	"github.com/occmundial/consumer-abe-atreel-user-message/sendgrid"
)

// NewConsumerService : Factory que crea un "ConsumerService"
func NewConsumerService(repository *repositories.ConsumerRepository) *ConsumerService {
	return &ConsumerService{
		Repository: repository,
	}
}

// ConsumerService :
type ConsumerService struct {
	Repository interfaces.IConsumerRepository
}

// IsOkProcessStart :
func (service *ConsumerService) IsOkProcessStart() models.StartStatus {
	success, err := service.Repository.IsHealthProcessToStart()
	return models.StartStatus{Success: success, Error: err}
}

// ProcessMessage :
func (service *ConsumerService) ProcessMessage() models.ProcessStatus {
	var (
		status string
		valid  bool
	)

	message, err := service.Repository.GetMessage()
	if err == nil {
		if valid, status, err = service.isValidMessage(message); valid {
			status, err = service.processAndCommitMessage(message)
			return models.ProcessStatus{Status: status, Message: message, Error: err}
		}
	} else {
		status = constants.StatusReadMessageError
	}
	go service.commitUnProccessMessage(message, status)
	return models.ProcessStatus{Status: status, Message: message, Error: err}
}

// isValidMessage :
// valid  => mensaje válido(true)/inválido(false)
// status => status para mensaje válido y procesado
// error  => error de procesamiento
func (service *ConsumerService) isValidMessage(message models.MessageForRead) (bool, string, error) {
	if !sendgrid.IsValidMessage(message.Message) {
		return false, constants.StatusAlreadyProcessed, nil
	}
	processed, err := service.Repository.IsProcessedMessage(message)
	if err != nil {
		// error durante verificación de mensaje
		return false, constants.StatusCheckProcessError, err
	}
	if processed {
		// el mensaje fue procesado con anterioridad o ya no se necesita procesarlo
		return false, constants.StatusAlreadyProcessed, nil
	}
	return true, "", nil
}

func (service *ConsumerService) processAndCommitMessage(message models.MessageForRead) (string, error) {
	status := ""
	err := service.processMessage(message)
	if err == nil {
		err = service.Repository.CommitMessage(message)
		if err == nil {
			status = constants.StatusFullProcessOK
		} else {
			// mensaje procesado exitosamente y no borrado de la cola
			status = constants.StatusProcessOK
		}
	} else {
		// Error en procesamiento del mensaje: falló el procesamiento del mensaje
		status = constants.StatusProcessError
	}

	return status, err
}

func (service *ConsumerService) processMessage(message models.MessageForRead) error {
	return service.Repository.CreateAndSendEmail(message.Message)
}

func (service *ConsumerService) commitUnProccessMessage(message models.MessageForRead, status string) {
	// es un evento válido de kafka Y
	// 		(se pudo leer el mensaje pero no se pudo deserializar O
	// 		es un mensaje inválido para lógica de negocio O
	// 		es un mensaje ya procesado)
	// => ya no se necesita procesar entonces lo damos como procesado
	if isValidEvent(message) && (status == constants.StatusReadMessageError ||
		status == constants.StatusInvalidMessage || status == constants.StatusAlreadyProcessed) {
		service.Repository.CommitMessage(message)
	}
}
