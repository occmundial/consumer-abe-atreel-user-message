package sendgrid

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/stretchr/testify/mock"
)

type mockSendgridRepository struct {
	mock.Mock
}

func (mockSR *mockSendgridRepository) GetDicState() (map[string]string, error) {
	args := mockSR.Called()
	return args.Get(0).(map[string]string), args.Error(1)
}

func (mockSR *mockSendgridRepository) CheckHealth() error {
	args := mockSR.Called()
	return args.Error(0)
}

func (mockSR *mockSendgridRepository) PostCorreo(messageFromKafka *models.MessageToProcess) error {
	args := mockSR.Called()
	return args.Error(0)
}
