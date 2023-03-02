package services

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/models"

	"github.com/stretchr/testify/mock"
)

type mockConsumerRepository struct {
	mock.Mock
}

// IsHealthProcessToStart :
func (mockCR *mockConsumerRepository) IsHealthProcessToStart() (bool, error) {
	args := mockCR.Called()
	return args.Bool(0), args.Error(1)
}

// GetMessage :
func (mockCR *mockConsumerRepository) GetMessage() (models.MessageForRead, error) {
	args := mockCR.Called()
	return args.Get(0).(models.MessageForRead), args.Error(1)
}

// CommitMessage :
func (mockCR *mockConsumerRepository) CommitMessage(message *models.MessageForRead) error {
	args := mockCR.Called(message)
	return args.Error(0)
}

// IsProcessedMessage :
func (mockCR *mockConsumerRepository) IsProcessedMessage(message *models.MessageForRead) (bool, error) {
	args := mockCR.Called(message)
	return args.Bool(0), args.Error(1)
}

func (mockCR *mockConsumerRepository) CreateAndSendEmail(message *models.MessageToProcess) error {
	args := mockCR.Called(&message)
	return args.Error(0)
}
