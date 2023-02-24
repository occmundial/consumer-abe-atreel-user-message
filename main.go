package main

import (
	"github.com/occmundial/consumer-abe-atreel-user-message/libs/logger"

	"github.com/occmundial/consumer-abe-atreel-user-message/di"
	"github.com/occmundial/consumer-abe-atreel-user-message/processor"
)

func main() {
	getWorker().Run()
}

func getWorker() *processor.Processor {
	var worker *processor.Processor

	container := di.GetContainer()
	err := container.Invoke(func(processor *processor.Processor) {
		worker = processor
	})

	if err != nil {
		logger.Fatal("main", "getWorker", err)
	}

	return worker
}
