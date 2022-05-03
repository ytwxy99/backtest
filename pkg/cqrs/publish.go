package cqrs

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/ytwxy99/backtest/pkg/database"
)

type PublishBus struct {
	Contract string
	Event    string
	Status   string
}

// Publish implements the method of the bus.EventBus interface.
func (publishBus *PublishBus) Publish(ctx context.Context) error {
	publish := database.Publish{
		Contract: publishBus.Contract,
		Event:    publishBus.Event,
		Status:   publishBus.Status,
	}

	err := publish.AddPublish(ctx)
	if err != nil {
		logrus.Error("add publish record faild: ", err)
		return err
	}

	return nil
}
