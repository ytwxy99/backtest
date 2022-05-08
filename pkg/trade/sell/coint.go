package sell

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/ytwxy99/backtest/pkg/database"
	"github.com/ytwxy99/backtest/pkg/utils"
	"github.com/ytwxy99/backtest/pkg/utils/market"
)

type CointSell struct {
	Contract  string
	Index     int
	Weights   []string
	Histories []*database.HistoryFourHour
}

// Sell sell the specified coin
func (cointSell *CointSell) Sell(ctx context.Context) error {
	average := &market.Average{
		CurrencyPair: cointSell.Contract,
		Level:        utils.Level4Hour,
		MA:           utils.MA21,
		Markets:      cointSell.Histories,
	}

	// for rising
	countRising := 0
	risingConditions := map[string]bool{}
	btcCurrentPrice := utils.StringToFloat64(cointSell.Histories[cointSell.Index].Price)
	if btcCurrentPrice < average.Average(false, cointSell.Index) {
		for _, weight := range cointSell.Weights {
			weightHistories, err := (&database.HistoryFourHour{
				Contract: weight,
			}).FetchHistoryFourHour(ctx)
			if err != nil {
				logrus.Error("fetch 4h history data failed: ", err)
				return err
			}

			if len(weightHistories) != len(cointSell.Histories) {
				logrus.Errorf("weight coins history data error: the length is %s, the coin is %s", len(weightHistories), weight)
				return err
			}

			averageWeight := &market.Average{
				CurrencyPair: weight,
				Level:        utils.Level4Hour,
				MA:           utils.MA21,
				Markets:      weightHistories,
			}

			weightCurrentPrice := utils.StringToFloat64(weightHistories[cointSell.Index].Price)
			if weightCurrentPrice < averageWeight.Average(false, cointSell.Index) {
				risingConditions[weight] = true
			}
		}

		// judge falling weight
		for _, condition := range risingConditions {
			if condition {
				countRising++
			}
		}

		if float32(countRising)/float32(len(cointSell.Weights)) >= 0.7 {
			// do sell
			order := &database.Order{
				Contract:  cointSell.Contract,
				Direction: utils.Up,
			}

			orders, err := order.FetchOrder(ctx)
			if err != nil {
				logrus.Error("fetch orders failed: ", err)
				return err
			}

			if len(orders) > 0 {
				if err := order.DeleteOrder(ctx); err != nil {
					logrus.Errorf("Sell order failed, the order detail is %s, the error is %s.", order, err)
					return err
				}
				logrus.Error(cointSell.Histories[cointSell.Index].Price)
			}
		}
	}

	// for falling
	countFalling := 0
	fallingConditions := map[string]bool{}
	btcCurrentPrice = utils.StringToFloat64(cointSell.Histories[cointSell.Index].Price)
	if btcCurrentPrice > average.Average(false, cointSell.Index) {
		for _, weight := range cointSell.Weights {
			weightHistories, err := (&database.HistoryFourHour{
				Contract: weight,
			}).FetchHistoryFourHour(ctx)
			if err != nil {
				logrus.Error("fetch 4h history data failed: ", err)
				return err
			}

			if len(weightHistories) != len(cointSell.Histories) {
				logrus.Errorf("weight coins history data error: the length is %s, the coin is %s", len(weightHistories), weight)
				return err
			}

			averageWeight := &market.Average{
				CurrencyPair: weight,
				Level:        utils.Level4Hour,
				MA:           utils.MA21,
				Markets:      weightHistories,
			}

			weightCurrentPrice := utils.StringToFloat64(weightHistories[cointSell.Index].Price)
			if weightCurrentPrice > averageWeight.Average(false, cointSell.Index) {
				fallingConditions[weight] = true
			}
		}

		// judge falling weight
		for _, condition := range risingConditions {
			if condition {
				countFalling++
			}
		}

		if float32(countFalling)/float32(len(cointSell.Weights)) >= 0.7 {
			// do sell
			order := &database.Order{
				Contract:  cointSell.Contract,
				Direction: utils.Down,
			}

			orders, err := order.FetchOrder(ctx)
			if err != nil {
				logrus.Error("fetch orders failed: ", err)
				return err
			}

			if len(orders) > 0 {
				if err := order.DeleteOrder(ctx); err != nil {
					logrus.Errorf("Sell order failed, the order detail is %s, the error is %s.", order, err)
					return err
				}
				logrus.Error(cointSell.Histories[cointSell.Index].Price)
			}
		}
	}

	return nil
}
