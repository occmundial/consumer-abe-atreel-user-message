package services

import (
	"strings"

	"github.com/occmundial/consumer-abe-atreel-user-message/models"
)

// isValidEvent: es un mensaje válido del framework de eventos
func isValidEvent(message models.MessageForRead) bool {
	return len(strings.TrimSpace(message.Key)) > 0
}
