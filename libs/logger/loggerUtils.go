package logger

import "github.com/occmundial/consumer-abe-atreel-user-message/models"

func isMessageWithValues(message models.MessageForRead) bool {
	return len(message.FlatMessage) > 0
}

func truncateText(text string) string {
	if len(text) > 100 {
		return text[:100] + "..."
	}
	return text
}
