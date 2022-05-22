package target

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/ytwxy99/autocoins/pkg/configuration"
	"github.com/ytwxy99/backtest/pkg/database"
	"github.com/ytwxy99/backtest/pkg/trade"
	"github.com/ytwxy99/backtest/pkg/trade/sell"
	"github.com/ytwxy99/backtest/pkg/utils"
	"github.com/ytwxy99/backtest/pkg/utils/market"
)

func (target *CointTarget) SearchTarget30M(ctx context.Context) map[string]interface{} {
	var err error
	histories, err := (&database.HistoryThirtyMin{
		Contract: target.TargetMetadata["contract"],
	}).FetchHistoryThirtyMin(ctx)

	if err != nil {
		logrus.Error("fetch 4h history data failed: ", err)
		return nil
	}

	weights, err := utils.ReadLines(ctx.Value("conf").(*configuration.SystemConf).WeightCsv)
	if err != nil {
		logrus.Error("Read weights coins failed: ", err)
		return nil
	}

	for i := 22; i < len(histories); i++ {
		// for rising market
		btcRisingCondition := conditionUpMonitor30M(target.TargetMetadata["contract"], 1.0015, i, histories)

		// for falling market
		btcFallingCondition := conditionDownMonitor30M(target.TargetMetadata["contract"], 1.0015, i, histories)

		// rising market buy point
		if btcRisingCondition {
			buyOrder := trade.BuyOrder{
				Contract:  histories[i].Contract,
				Price:     histories[i].Price,
				Direction: utils.Up,
				BuyTime:   histories[i].Time,
			}
			if err := buyOrder.Buy(ctx); err != nil {
				logrus.Error("coint back test terminal, the reason is ", err)
			}
		}

		// falling market buy point
		if btcFallingCondition {
			buyOrder := trade.BuyOrder{
				Contract:  histories[i].Contract,
				Price:     histories[i].Price,
				Direction: utils.Down,
				BuyTime:   histories[i].Time,
			}
			if err := buyOrder.Buy(ctx); err != nil {
				logrus.Error("coint back test terminal, the reason is ", err)
			}
		}

		// judge sell point
		err = (&sell.CointSell{
			Contract:  histories[i].Contract,
			Level: utils.Level30Min,
			Index:     i,
			Weights:   weights,
			Histories: histories,
		}).Sell(ctx)

		if err != nil {
			return map[string]interface{}{
				target.TargetMetadata["contract"]: 0,
			}
		}
	}

	return nil
}

func conditionUpMonitor30M(coin string, tenAverageDiff float64, index int, markets []*database.HistoryThirtyMin) bool {
	averageArgs := &market.Average{
		CurrencyPair: coin,
		Level:        utils.Level30Min,
		MA:           utils.MA21,
		Markets:      markets,
	}
	MA21Average := averageArgs.Average(false, index) > averageArgs.Average(true, index)*tenAverageDiff //4h的Average是增长的

	averageArgs.MA = utils.MA5
	MA5Average := averageArgs.Average(false, index) > averageArgs.Average(true, index)

	priceC := utils.StringToFloat64(averageArgs.Markets.([]*database.HistoryThirtyMin)[index].Price) > averageArgs.Average(false, index)

	return MA21Average && MA5Average && priceC
}

func conditionDownMonitor30M(coin string, averageDiff float64, index int, markets []*database.HistoryThirtyMin) bool {
	averageArgs := &market.Average{
		CurrencyPair: coin,
		Level:        utils.Level30Min,
		MA:           utils.MA21,
		Markets:      markets,
	}
	MA21Average := averageArgs.Average(false, index)*averageDiff < averageArgs.Average(true, index) //4h的Average是增长的

	averageArgs.MA = utils.MA5
	MA5Average := averageArgs.Average(false, index) < averageArgs.Average(true, index)

	priceC := utils.StringToFloat64(averageArgs.Markets.([]*database.HistoryThirtyMin)[index].Price) < averageArgs.Average(false, index)

	return MA21Average && MA5Average && priceC
}