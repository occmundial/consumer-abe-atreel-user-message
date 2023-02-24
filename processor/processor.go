package processor

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/occmundial/consumer-abe-atreel-user-message/services"
)

func NewProcessor(configuration *models.Configuration, service *services.ConsumerService) *Processor {
	return &Processor{
		Configuration: configuration,
		Service:       service,
	}
}

type Processor struct {
	Configuration *models.Configuration
	Service       interfaces.IConsumerService
}

func (processor *Processor) Run() {
	for {
		if processor.isOkProcessStart() {
			for {
				if !processor.processMessage() {
					break
				}
			}
		} else {
			logger.LogStatusDebug(constants.StatusCheckProcessError)
		}
		wait(processor.Configuration.QueueRequestDelay)
	}
}

func (processor *Processor) isOkProcessStart() bool {
	status := processor.Service.IsOkProcessStart()
	if status.Error != nil {
		logger.LogError("processor", "isOkProcessStart", status.Error.Error())
	}
	return status.Success
}

func (processor *Processor) processMessage() bool {
	logger.LogSimpleDebug(fmt.Sprintf("process started (version: %s)", processor.Configuration.ArtifactVersion))
	processStatus := processor.Service.ProcessMessage()
	logger.LogSimpleDebug(fmt.Sprintf("process ended (version: %s)", processor.Configuration.ArtifactVersion))
	logStatus(processStatus)
	return isHealthyStatus(processStatus)
}

func logStatus(processStatus models.ProcessStatus) {
	isError, value := getError(processStatus.Error)
	logger.LogConditionalStatus(processStatus, isError, value)
}

func getError(err error) (bool, string) {
	if err != nil {
		return !errors.Is(err, context.DeadlineExceeded), err.Error()
	}
	return false, ""
}

func wait(delay int64) {
	logger.LogSimpleDebug(fmt.Sprintf("wait for %d seconds...", delay))
	time.Sleep(time.Duration(delay) * time.Second)
}
