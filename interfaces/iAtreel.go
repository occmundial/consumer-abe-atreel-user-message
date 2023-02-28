package interfaces

import "github.com/occmundial/consumer-abe-atreel-user-message/models"

type IAtreel interface {
	PostCorreo(messageFromKafka *models.MessageToProcess) error
}
