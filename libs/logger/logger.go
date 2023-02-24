package logger

import (
	"fmt"
	"os"

	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
}

func LogSimpleDebug(text string) {
	log.Debug().Msg(text)
}

func LogStatusDebug(status string) {
	log.Debug().
		Str("status", status).
		Msg("")
}

func LogConditionalStatus(processStatus models.ProcessStatus, isError bool, info string) {
	var logConditional *zerolog.Event
	var fieldName = "info"
	if !isError {
		logConditional = log.Info()
	} else {
		logConditional = log.Error()
		fieldName = "error"
	}
	if isMessageWithValues(processStatus.Message) {
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

func LogError(moduleName string, functionName string, errorText string) {
	log.Error().
		Str("moduleName", moduleName).
		Str("functionName", functionName).
		Str("error", errorText).
		Msg("")
}

func LogInfo(moduleName string, functionName string) {
	log.Info().
		Str("moduleName", moduleName).
		Str("functionName", functionName).
		Msg("")
}

// ConditionalFatal :
func ConditionalFatal(moduleName string, functionName string, errs ...error) {
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
func Fatal(moduleName string, functionName string, err error) {
	FatalS(moduleName, functionName, err.Error())
}

// Fatal :
func FatalS(moduleName string, functionName string, err string) {
	log.Fatal().
		Str("moduleName", moduleName).
		Str("functionName", functionName).
		Str("error", err).
		Msg("")
}
