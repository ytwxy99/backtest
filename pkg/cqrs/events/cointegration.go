package events

import (
	"context"

	"github.com/ytwxy99/autocoins/pkg/configuration"
	"github.com/ytwxy99/backtest/pkg/database"

	"github.com/ytwxy99/backtest/pkg/trade/target"
	"github.com/ytwxy99/backtest/pkg/utils"
)

type CointegrationEvent struct {
	EventMetadata map[string]string
}

func (cointegrationEvent *CointegrationEvent) DoEvent(ctx context.Context) (string, error) {
	btcTarget := &target.CointBtcTarget{
		TargetMetadata: cointegrationEvent.EventMetadata,
	}

	if btcTarget.TargetMetadata["contract"] == utils.ALL {
		if err := (&database.Order{}).DeleteALLLOrder(ctx); err != nil {
			return "coint failed", err
		}

		coins, err := utils.ReadLines(ctx.Value("conf").(*configuration.SystemConf).WeightCsv)
		if err != nil {
			return "coint failed", err
		}

		for _, coin := range coins {
			btcTarget.TargetMetadata["contract"] = coin
			btcTarget.SearchTarget(ctx)
		}
	} else {
		btcTarget.SearchTarget(ctx)
	}

	return "coint", nil
}
