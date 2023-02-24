package interfaces

import "github.com/occmundial/consumer-abe-atreel-user-message/models"

// IConsumerService :
type IConsumerService interface {
	IsOkProcessStart() models.StartStatus
	ProcessMessage() models.ProcessStatus
}
