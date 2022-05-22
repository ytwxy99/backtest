package cqrs

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/ytwxy99/backtest/pkg/cqrs/events"
	"github.com/ytwxy99/backtest/pkg/utils"
)

var event events.Event

type AsynchronousDispatchMetadata struct {
	Metadata map[string]string
}

func (asynchronousDispatchMetadata *AsynchronousDispatchMetadata) Dispatch(ctx context.Context) error {
	logrus.Info("Fetch a new event and dispatch it : ", asynchronousDispatchMetadata.Metadata)
	metadata := asynchronousDispatchMetadata.Metadata
	if metadata["event"] == utils.Coint4h {
		event = &events.Coin4hEvent{
			EventMetadata: metadata,
		}
	} else if metadata["event"] == utils.Coint30m {
		event = &events.Coin30mEvent{
			EventMetadata: metadata,
		}
	} else if metadata["event"] == utils.Iint {
		event = &events.InitEvent{
			EventMetadata: metadata,
		}
	}

	response, err := event.DoEvent(ctx)
	if err != nil {
		logrus.Error("Do event failed: ", err)
		ErrResponse <- err
		return err
	}

	asynchronousDispatchMetadata.Metadata["response"] = response
	DispatchResponse <- asynchronousDispatchMetadata.Metadata
	return nil
}
