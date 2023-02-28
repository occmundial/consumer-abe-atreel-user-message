package processor

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

// isHealthyStatus : sin error y la cola tiene al menos un mensaje y kafka está saludable
func isHealthyStatus(status *models.ProcessStatus) bool {
	return status.Error == nil && status.Status != constants.StatusProcessStartError
}
