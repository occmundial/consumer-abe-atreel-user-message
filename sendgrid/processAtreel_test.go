package sendgrid

import (
	"errors"
	"testing"

	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	messageToProcess = models.MessageToProcess{
		Email: "email@test.com",
	}
	config = models.Configuration{}
)

func Test_ProcessAtreel_IsValidMessage_Fail(t *testing.T) {
	valid := IsValidMessage(&models.MessageToProcess{})
	assert.False(t, valid)
}

func Test_ProcessAtreel_IsValidMessage_OK(t *testing.T) {
	valid := IsValidMessage(&models.MessageToProcess{Email: "test@domain.com"})
	assert.True(t, valid)
}

func Test_CreateAndSendEmail_Init_OK(t *testing.T) {
	mockRepository := new(mockSendgridRepository)
	mockRepository.On("GetDicState").Return(make(map[string]string), nil)
	processor := AtreelProcessor{Configuration: &config, Atreel: mockRepository}
	assert.NotNil(t, processor)
}

func Test_CreateAndSendEmail_OK(t *testing.T) {
	mockRepository := new(mockSendgridRepository)
	mockRepository.On("GetDicState").Return(make(map[string]string), nil)
	mockRepository.On("PostCorreo", mock.Anything).Return(nil)
	processor := AtreelProcessor{Configuration: &config, Atreel: mockRepository}
	err := processor.CreateAndSendEmail(&messageToProcess)
	assert.NoError(t, err)
}

func Test_CreateAndSendEmail_Fail(t *testing.T) {
	fakeError := errors.New("err fake")
	mockRepository := new(mockSendgridRepository)
	mockRepository.On("GetDicState").Return(make(map[string]string), nil)
	mockRepository.On("PostCorreo", mock.Anything).Return(fakeError)
	processor := AtreelProcessor{Configuration: &config, Atreel: mockRepository}
	err := processor.CreateAndSendEmail(&messageToProcess)
	assert.Error(t, err)
	assert.Equal(t, fakeError, err)
}
