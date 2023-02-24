package sendgrid

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/stretchr/testify/mock"
)

type mockSendgridRepository struct {
	mock.Mock
}

func (mock *mockSendgridRepository) GetDicState() (map[string]string, error) {
	args := mock.Called()
	return args.Get(0).(map[string]string), args.Error(1)
}

func (mock *mockSendgridRepository) CheckHealth() error {
	args := mock.Called()
	return args.Error(0)
}

func (mock *mockSendgridRepository) PostCorreo(messageFromKafka models.MessageToProcess) error {
	args := mock.Called()
	return args.Error(0)
}
