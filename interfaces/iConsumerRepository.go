package interfaces

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

// IConsumerRepository :
type IConsumerRepository interface {
	IsHealthProcessToStart() (bool, error)
	GetMessage() (models.MessageForRead, error)
	IsProcessedMessage(message *models.MessageForRead) (bool, error)
	CommitMessage(message *models.MessageForRead) error
	CreateAndSendEmail(msg *models.MessageToProcess) error
}
