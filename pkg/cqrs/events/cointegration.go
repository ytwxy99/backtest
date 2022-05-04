package events

import (
	"context"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/ytwxy99/backtest/pkg/database"
)

type CointegrationEvent struct {
	EventMetadata map[string]string
}

func (cointegrationEvent *CointegrationEvent) DoEvent(ctx context.Context) (string, error) {
	history := &database.HistoryDay{
		Contract: cointegrationEvent.EventMetadata["contract"],
	}
	histories, err := history.FetchHistoryDay(ctx)
	if err != nil {
		logrus.Error("Fetach history data failed: ", err)
	}
	fmt.Println(len(histories))

	return "", nil
}
