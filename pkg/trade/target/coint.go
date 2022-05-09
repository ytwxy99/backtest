package target

import (
	"context"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/ytwxy99/autocoins/pkg/configuration"

	"github.com/ytwxy99/backtest/pkg/database"
	"github.com/ytwxy99/backtest/pkg/trade"
	"github.com/ytwxy99/backtest/pkg/trade/sell"
	"github.com/ytwxy99/backtest/pkg/utils"
	"github.com/ytwxy99/backtest/pkg/utils/market"
)

type CointBtcTarget struct {
	TargetMetadata map[string]string
}

func (target *CointBtcTarget) SearchTarget(ctx context.Context) map[string]interface{} {
	histories, err := (&database.HistoryFourHour{
		Contract: target.TargetMetadata["contract"],
	}).FetchHistoryFourHour(ctx)
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
		btcRisingCondition := conditionUpMonitor(target.TargetMetadata["contract"], 1.0015, i, histories)

		// for falling market
		btcFallingCondition := conditionDownMonitor(target.TargetMetadata["contract"], 1.0015, i, histories)

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

func conditionUpMonitor(coin string, tenAverageDiff float64, index int, markets []*database.HistoryFourHour) bool {
	averageArgs := &market.Average{
		CurrencyPair: coin,
		Level:        utils.Level4Hour,
		MA:           utils.MA21,
		Markets:      markets,
	}
	MA21Average := averageArgs.Average(false, index) > averageArgs.Average(true, index)*tenAverageDiff //4h的Average是增长的

	averageArgs.MA = utils.MA5
	MA5Average := averageArgs.Average(false, index) > averageArgs.Average(true, index)

	priceC := utils.StringToFloat64(averageArgs.Markets.([]*database.HistoryFourHour)[index].Price) > averageArgs.Average(false, index)

	return MA21Average && MA5Average && priceC
}

func conditionDownMonitor(coin string, averageDiff float64, index int, markets []*database.HistoryFourHour) bool {
	averageArgs := &market.Average{
		CurrencyPair: coin,
		Level:        utils.Level4Hour,
		MA:           utils.MA21,
		Markets:      markets,
	}
	MA21Average := averageArgs.Average(false, index)*averageDiff < averageArgs.Average(true, index) //4h的Average是增长的

	averageArgs.MA = utils.MA5
	MA5Average := averageArgs.Average(false, index) < averageArgs.Average(true, index)

	priceC := utils.StringToFloat64(averageArgs.Markets.([]*database.HistoryFourHour)[index].Price) < averageArgs.Average(false, index)

	return MA21Average && MA5Average && priceC
}

type cSort struct {
	Pair  string
	Value float32
}

func (*cSort) sortCoints(coints map[string]float32) []cSort {
	var cSorts []cSort
	for k, v := range coints {
		cSorts = append(cSorts, cSort{k, v})
	}

	sort.Slice(cSorts, func(i, j int) bool {
		return cSorts[i].Value < cSorts[j].Value // 升序
	})

	return cSorts
}

func averageDiff(coin string, level string, index int, markets []*database.HistoryFourHour) bool {
	var maValues []float64
	var max float64
	var min float64

	mas := []int{
		utils.MA5,
		utils.MA10,
		utils.MA21,
	}

	for _, ma := range mas {
		averageArgs := &market.Average{
			CurrencyPair: coin,
			Level:        level,
			MA:           ma,
			Markets:      markets,
		}
		maValues = append(maValues, averageArgs.Average(false, index))
	}

	if len(maValues) != 3 {
		return false
	}

	for _, value := range maValues {
		if max == 0 {
			max = value
			min = value
		}

		if value >= max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	return (max-min)/min > 0.01
}
