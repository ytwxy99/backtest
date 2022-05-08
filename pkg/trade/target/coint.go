package target

import (
	"context"
	"sort"

	"github.com/sirupsen/logrus"
	"github.com/ytwxy99/autocoins/pkg/configuration"

	"github.com/ytwxy99/backtest/pkg/trade"
	"github.com/ytwxy99/backtest/pkg/database"
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
		conditionsRising := map[string]bool{}
		conditionsFalling := map[string]bool{}

		average := &market.Average{
			CurrencyPair: target.TargetMetadata["contract"],
			Level:        utils.Level4Hour,
			MA:           utils.MA21,
			Markets:      histories,
		}

		// for rising market
		btcRisingCondition := conditionUpMonitor(target.TargetMetadata["contract"], 1.0, i, histories)
		priceRisingCondition := utils.StringToFloat64(histories[i].Price) > average.Average(false, i)

		for _, weight := range weights {
			weight_histories, err := (&database.HistoryFourHour{
				Contract: weight,
			}).FetchHistoryFourHour(ctx)
			if err != nil {
				logrus.Error("fetch 4h history data failed: ", err)
				return nil
			}

			if len(weight_histories) != len(histories) {
				logrus.Errorf("weight coins history data error: the length is %s, the btc length is %s, the coin is %s", len(weight_histories), len(histories), weight)
				return nil
			}

			conditionsRising[weight] = conditionUpMonitor(weight, 1.0, i, weight_histories)
		}

		// for falling market
		btcFallingCondition := conditionDownMonitor(target.TargetMetadata["contract"], 1.0, i, histories)
		priceFallingCondition := utils.StringToFloat64(histories[i].Price) < average.Average(false, i)

		for _, weight := range weights {
			weight_histories, err := (&database.HistoryFourHour{
				Contract: weight,
			}).FetchHistoryFourHour(ctx)
			if err != nil {
				logrus.Error("fetch 4h history data failed: ", err)
				return nil
			}

			if len(weight_histories) != len(histories) {
				logrus.Errorf("weight coins history data error: the length is %s, the coin is %s", len(weight_histories), weight)
				return nil
			}

			conditionsFalling[weight] = conditionDownMonitor(weight, 1.0, i, weight_histories)
		}

		// judge rising weight
		countUp := 0
		allUp := 0
		for _, condition := range conditionsRising {
			if condition {
				countUp++
			}
			allUp++
		}

		// judge falling weight
		countDown := 0
		allDown := 0
		for _, condition := range conditionsFalling {
			if condition {
				countDown++
			}
			allDown++
		}

		// rising market buy point
		if float32(countUp)/float32(allUp) >= 0.7 && btcRisingCondition && priceRisingCondition && averageDiff(target.TargetMetadata["contract"], utils.Level4Hour, i, histories) {

			buyOrder := trade.BuyOrder{
				Contract:  histories[i].Contract,
				Price:     histories[i].Price,
				Direction: utils.Up,
			}
			if err := buyOrder.Buy(ctx); err != nil {
				logrus.Error("coint back test terminal, the reason is ", err)
			}
		}

		if float32(countDown)/float32(allDown) >= 0.7 && btcFallingCondition && priceFallingCondition && averageDiff(target.TargetMetadata["contract"], utils.Level4Hour, i, histories) {
			buyOrder := trade.BuyOrder{
				Contract:  histories[i].Contract,
				Price:     histories[i].Price,
				Direction: utils.Down,
			}
			if err := buyOrder.Buy(ctx); err != nil {
				logrus.Error("coint back test terminal, the reason is ", err)
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

	averageArgs.MA = utils.MA10
	MA10Average := averageArgs.Average(false, index) > averageArgs.Average(true, index)

	averageArgs.MA = utils.MA5
	MA15Average := averageArgs.Average(false, index) > averageArgs.Average(true, index)

	return MA21Average && MA10Average && MA15Average
}

func conditionDownMonitor(coin string, averageDiff float64, index int, markets []*database.HistoryFourHour) bool {
	averageArgs := &market.Average{
		CurrencyPair: coin,
		Level:        utils.Level4Hour,
		MA:           utils.MA21,
		Markets:      markets,
	}
	MA21Average := averageArgs.Average(false, index)*averageDiff < averageArgs.Average(true, index) //4h的Average是增长的

	averageArgs.MA = utils.MA10
	MA10Average := averageArgs.Average(false, index) < averageArgs.Average(true, index)

	averageArgs.MA = utils.MA5
	MA15Average := averageArgs.Average(false, index) < averageArgs.Average(true, index)

	return MA21Average && MA10Average && MA15Average
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

	return (max-min)/min > 0.03
}
