package logger

import "github.com/occmundial/consumer-abe-atreel-user-message/models"

const MaxCharactersByLine = 100

func isMessageWithValues(message *models.MessageForRead) bool {
	return len(message.FlatMessage) > 0
}

func truncateText(text string) string {
	if len(text) > MaxCharactersByLine {
		return text[:MaxCharactersByLine] + "..."
	}
	return text
}
