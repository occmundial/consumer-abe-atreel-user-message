package services

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/models"

	"github.com/stretchr/testify/mock"
)

type mockConsumerRepository struct {
	mock.Mock
}

// IsOkProcessStart :
func (mock *mockConsumerRepository) IsHealthProcessToStart() (bool, error) {
	args := mock.Called()
	return args.Bool(0), args.Error(1)
}

// GetMessage :
func (mock *mockConsumerRepository) GetMessage() (models.MessageForRead, error) {
	args := mock.Called()
	return args.Get(0).(models.MessageForRead), args.Error(1)
}

// CommitMessage :
func (mock *mockConsumerRepository) CommitMessage(message models.MessageForRead) error {
	args := mock.Called(message)
	return args.Error(0)
}

// IsProcessedMessage :
func (mock *mockConsumerRepository) IsProcessedMessage(message models.MessageForRead) (bool, error) {
	args := mock.Called(message)
	return args.Bool(0), args.Error(1)
}

func (mock *mockConsumerRepository) CreateAndSendEmail(message models.MessageToProcess) error {
	args := mock.Called(message)
	return args.Error(0)
}
