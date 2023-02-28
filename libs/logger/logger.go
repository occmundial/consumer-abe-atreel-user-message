package logger

import (
	"fmt"

	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LogSimpleDebug(text string) {
	log.Debug().Msg(text)
}

func LogStatusDebug(status string) {
	log.Debug().
		Str("status", status).
		Msg("")
}

func LogConditionalStatus(processStatus *models.ProcessStatus, isError bool, info string) {
	var logConditional *zerolog.Event
	var fieldName = "info"
	if !isError {
		logConditional = log.Info()
	} else {
		logConditional = log.Error()
		fieldName = "error"
	}
	if isMessageWithValues(&processStatus.Message) {
		logConditional.
			Str("status", processStatus.Status).
			Str("topic", processStatus.Message.Topic).
			Str("key", processStatus.Message.Key).
			Str("message", truncateText(processStatus.Message.FlatMessage)).
			Str(fieldName, info).
			Msg("")
	} else {
		logConditional.
			Str("status", processStatus.Status).
			Str(fieldName, info).
			Msg("")
	}
}

func LogError(moduleName, functionName, errorText string) {
	log.Error().
		Str("moduleName", moduleName).
		Str("functionName", functionName).
		Str("error", errorText).
		Msg("")
}

func LogInfo(moduleName, functionName string) {
	log.Info().
		Str("moduleName", moduleName).
		Str("functionName", functionName).
		Msg("")
}

// ConditionalFatal :
func ConditionalFatal(moduleName, functionName string, errs ...error) {
	text := ""
	for _, err := range errs {
		if err != nil {
			text += fmt.Sprintf("%s \t", err.Error())
		}
	}
	if len(text) > 0 {
		FatalS(moduleName, functionName, text)
	}
}

// Fatal :
func Fatal(moduleName, functionName string, err error) {
	FatalS(moduleName, functionName, err.Error())
}

// Fatal :
func FatalS(moduleName, functionName, err string) {
	log.Fatal().
		Str("moduleName", moduleName).
		Str("functionName", functionName).
		Str("error", err).
		Msg("")
}
