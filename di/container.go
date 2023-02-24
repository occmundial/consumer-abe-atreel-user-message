package di

import (
	"sync"

	"github.com/occmundial/consumer-abe-atreel-user-message/config"
	"github.com/occmundial/consumer-abe-atreel-user-message/database"
	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"
	"github.com/occmundial/consumer-abe-atreel-user-message/processor"
	"github.com/occmundial/consumer-abe-atreel-user-message/repositories"
	"github.com/occmundial/consumer-abe-atreel-user-message/requests"
	"github.com/occmundial/consumer-abe-atreel-user-message/sendgrid"
	"github.com/occmundial/consumer-abe-atreel-user-message/services"
	"go.uber.org/dig"
)

// https://blog.drewolson.org/dependency-injection-in-go

var (
	container *dig.Container
	once      sync.Once
)

// GetContainer :
func GetContainer() *dig.Container {
	once.Do(func() {
		container = buildContainer()
	})
	return container
}

// buildContainer :
func buildContainer() *dig.Container {
	c := dig.New()
	logger.ConditionalFatal("container", "buildContainer",
		c.Provide(config.NewConfiguration),
		c.Provide(database.NewSQLServer),
		c.Provide(requests.NewRequests),
		c.Provide(database.NewQueries),
		c.Provide(requests.NewAtreel),
		c.Provide(sendgrid.NewAtreelProcessor),
		c.Provide(repositories.NewConsumerRepository),
		c.Provide(services.NewConsumerService),
		c.Provide(processor.NewProcessor))
	return c
}
