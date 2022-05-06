package events

import (
	"context"

	"github.com/ytwxy99/backtest/pkg/trade/target"
	"github.com/ytwxy99/backtest/pkg/utils"
)

type CointegrationEvent struct {
	EventMetadata map[string]string
}

func (cointegrationEvent *CointegrationEvent) DoEvent(ctx context.Context) (string, error) {
	if cointegrationEvent.EventMetadata["contract"] == utils.BTC {
		btcTarget := &target.CointBtcTarget{
			TargetMetadata: cointegrationEvent.EventMetadata,
		}
		btcTarget.SearchTarget(ctx)
	}

	return "coint", nil
}
