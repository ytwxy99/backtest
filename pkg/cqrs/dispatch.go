package cqrs

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/ytwxy99/backtest/pkg/cqrs/events"
	"github.com/ytwxy99/backtest/pkg/utils"
	"strings"
)

var event events.Event

type AsynchronousDispatchMetadata struct {
	Metadata map[string]string
}

func (asynchronousDispatchMetadata *AsynchronousDispatchMetadata) Dispatch(ctx context.Context) error {
	logrus.Info("Fetch a new event and dispatch it : ", asynchronousDispatchMetadata.Metadata)
	//todo(wangxiaoyu), implement detail
	if strings.Contains(asynchronousDispatchMetadata.Metadata["event"], utils.Cointegration) {
		event = &events.CointegrationEvent{
			EventMetadata: asynchronousDispatchMetadata.Metadata,
		}
	}

	response, err := event.DoEvent(ctx)
	if err != nil {
		logrus.Error("Do event failed: ", err)
	}

	asynchronousDispatchMetadata.Metadata["response"] = response
	DispatchResponse <- asynchronousDispatchMetadata.Metadata
	return nil
}
