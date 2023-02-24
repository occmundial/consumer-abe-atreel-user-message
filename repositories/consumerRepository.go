package repositories

import (
	"context"
	"time"

	"github.com/occmundial/consumer-abe-atreel-user-message/database"
	"github.com/occmundial/consumer-abe-atreel-user-message/interfaces"
	"github.com/occmundial/consumer-abe-atreel-user-message/models"
	"github.com/occmundial/consumer-abe-atreel-user-message/sendgrid"

	"github.com/segmentio/kafka-go"
)

var (
	reader       *kafka.Reader
	kafkaTimeout time.Duration
)

// NewConsumerRepository : Factory que crea un "ConsumerRepository"
func NewConsumerRepository(configuration *models.Configuration, consumerSengrid *sendgrid.AtreelProcessor, queries *database.Queries) *ConsumerRepository {
	cr := ConsumerRepository{Configuration: configuration, AtreelProcessor: consumerSengrid, Queries: queries}
	cr.init()
	return &cr
}

// ConsumerRepository :
type ConsumerRepository struct {
	Configuration   *models.Configuration
	AtreelProcessor *sendgrid.AtreelProcessor
	Queries         interfaces.IQuery
}

func (repository ConsumerRepository) init() {
	kafkaTimeout = time.Duration(repository.Configuration.QueueTimeout) * time.Second
	reader = getKafkaReader(repository.Configuration)
}

func (ConsumerRepository) GetMessage() (models.MessageForRead, error) {
	ctx, cancel := context.WithTimeout(context.Background(), kafkaTimeout)
	defer cancel()
	kafkaMessage, err := reader.FetchMessage(ctx)
	if err != nil {
		return models.MessageForRead{}, err
	}
	return kafkaMessageToMessageForRead(kafkaMessage)
}

func (ConsumerRepository) CommitMessage(message models.MessageForRead) error {
	ctx, cancel := context.WithTimeout(context.Background(), kafkaTimeout)
	defer cancel()
	return reader.CommitMessages(ctx, message.SourceMessage)
}

func (repository ConsumerRepository) IsHealthProcessToStart() (bool, error) {
	chanHealth := make(chan string)
	defer closeChannels(chanHealth)
	go processHealth(chanHealth, repository.Configuration, repository.Queries)
	emailsHealth := <-chanHealth
	err := concatErrors(emailsHealth)
	return err == nil, err
}

// IsProcessedMessage :
func (repository ConsumerRepository) IsProcessedMessage(message models.MessageForRead) (bool, error) {
	return false, nil
}

func (repository ConsumerRepository) CreateAndSendEmail(message models.MessageToProcess) error {
	return repository.AtreelProcessor.CreateAndSendEmail(message)
}
