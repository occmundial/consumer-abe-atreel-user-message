package services

import (
	"errors"
	"testing"

	"github.com/occmundial/consumer-abe-atreel-user-message/constants"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	messageForRead = &models.MessageForRead{
		Key:   "c4bf18e8-10f6-48ca-aa6f-61c7e8c25578",
		Topic: "test-successful",
		Message: models.MessageToProcess{
			Email: "email@test.com",
		},
	}
)

// todo perfecto => happy path (created event)
func TestConsumerService_When_MessageIsOK_Expect_StatusProcessOK(t *testing.T) {
	mockRepository := new(mockConsumerRepository)

	mockRepository.On("IsOkProcessStart").Return(true, nil)
	mockRepository.On("GetMessage").Return(*messageForRead, nil)
	mockRepository.On("IsProcessedMessage", messageForRead).Return(false, nil)
	mockRepository.On("CommitMessage", messageForRead).Return(nil)
	mockRepository.On("CreateAndSendEmail", mock.Anything).Return(nil)

	expectedProcessStatus := models.ProcessStatus{
		Status:  constants.StatusFullProcessOK,
		Message: *messageForRead,
	}

	consumerService := ConsumerService{Repository: mockRepository}
	actualProcessStatus := consumerService.ProcessMessage()

	assert.Equal(t, expectedProcessStatus.Message.Key, actualProcessStatus.Message.Key)
	assert.Equal(t, expectedProcessStatus.Status, actualProcessStatus.Status)
	mockRepository.AssertCalled(t, "CommitMessage", messageForRead)
}

// error en revisión del mensaje a procesar
func TestConsumerService_When_CheckProcessError_Expect_StatusCheckProcessError(t *testing.T) {
	mockRepository := new(mockConsumerRepository)
	err := errors.New("error en verificación de procesamiento")
	mockRepository.On("IsHealthProcessToStart").Return(true, nil)
	mockRepository.On("GetMessage").Return(*messageForRead, nil)
	mockRepository.On("IsProcessedMessage", messageForRead).Return(false, err)
	mockRepository.On("CreateAndSendEmail", mock.Anything).Return(nil)

	expectedProcessStatus := models.ProcessStatus{
		Status:  constants.StatusCheckProcessError,
		Message: *messageForRead,
		Error:   err,
	}

	consumerService := ConsumerService{Repository: mockRepository}
	actualProcessStatus := consumerService.ProcessMessage()

	assert.Equal(t, expectedProcessStatus.Status, actualProcessStatus.Status)
}

// el mensaje ya había sido procesado
func TestConsumerService_When_MessageProcessed_Expect_StatusNoProcessMessage(t *testing.T) {
	mockRepository := new(mockConsumerRepository)
	mockRepository.On("IsOkProcessStart").Return(true, nil)
	mockRepository.On("GetMessage").Return(*messageForRead, nil)
	mockRepository.On("IsProcessedMessage", messageForRead).Return(true, nil)
	mockRepository.On("CommitMessage", messageForRead).Return(nil)

	expectedProcessStatus := models.ProcessStatus{
		Status:  constants.StatusAlreadyProcessed,
		Message: *messageForRead,
	}

	consumerService := ConsumerService{Repository: mockRepository}
	actualProcessStatus := consumerService.ProcessMessage()

	assert.Equal(t, expectedProcessStatus.Status, actualProcessStatus.Status)
}

// mensaje procesado exitosamente pero fallo en commit del mensaje
func TestConsumerService_When_CommitMessageError_Expect_StatusProcessError(t *testing.T) {
	mockRepository := new(mockConsumerRepository)
	err := errors.New("error en commit (borrado) del mensaje")
	mockRepository.On("CommitMessage", mock.Anything).Return(err)
	mockRepository.On("IsOkProcessStart").Return(true, nil)
	mockRepository.On("GetMessage").Return(*messageForRead, nil)
	mockRepository.On("IsProcessedMessage", messageForRead).Return(false, nil)

	mockRepository.On("CreateAndSendEmail", mock.Anything).Return(nil)
	expectedProcessStatus := models.ProcessStatus{
		Status:  constants.StatusProcessOK,
		Message: *messageForRead,
		Error:   err,
	}
	consumerService := ConsumerService{Repository: mockRepository}
	actualProcessStatus := consumerService.ProcessMessage()
	assert.Equal(t, expectedProcessStatus.Status, actualProcessStatus.Status)
	mockRepository.AssertCalled(t, "CommitMessage", messageForRead)
}

func TestConsumerService_When_MessageNotGetMessage_Expect_Fail(t *testing.T) {
	mockRepository := new(mockConsumerRepository)
	mockRepository.On("IsOkProcessStart").Return(true, nil)
	mockRepository.On("GetMessage").Return(*messageForRead, errors.New(constants.StatusReadMessageError))
	mockRepository.On("IsProcessedMessage", messageForRead).Return(false, nil)
	mockRepository.On("CommitMessage", mock.Anything).Return(nil)
	consumerService := ConsumerService{Repository: mockRepository}
	actualProcessStatus := consumerService.ProcessMessage()
	assert.Equal(t, actualProcessStatus.Error.Error(), constants.StatusReadMessageError)
}

func TestConsumerService_When_InvalidMessage_Expect_Fail(t *testing.T) {
	mockRepository := new(mockConsumerRepository)
	messageForRead.Message.Email = ""
	mockRepository.On("IsOkProcessStart").Return(true, nil)
	mockRepository.On("GetMessage").Return(*messageForRead, nil)
	mockRepository.On("IsProcessedMessage", messageForRead).Return(false, nil)
	mockRepository.On("CommitMessage", messageForRead).Return(nil)
	expectedProcessStatus := models.ProcessStatus{
		Status:  constants.StatusAlreadyProcessed,
		Message: *messageForRead,
	}
	consumerService := ConsumerService{Repository: mockRepository}
	actualProcessStatus := consumerService.ProcessMessage()
	assert.Equal(t, expectedProcessStatus.Message.Key, actualProcessStatus.Message.Key)
	assert.Equal(t, expectedProcessStatus.Status, actualProcessStatus.Status)
}
