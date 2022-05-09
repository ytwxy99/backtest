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
	orders, err := (&database.Order{
		Contract:  cointSell.Contract,
		Direction: utils.Up,
	}).FetchOrder(ctx)
	if err != nil {
		logrus.Error("fetch orders failed: ", err)
		return err
	}

	r1 := average.Average(false, cointSell.Index) <= average.Average(true, cointSell.Index)
	r2 := len(orders) == 1 && utils.StringToFloat64(cointSell.Histories[cointSell.Index].Price)*1.01 <= utils.StringToFloat64(orders[0].Price)
	r3 := utils.StringToFloat64(cointSell.Histories[cointSell.Index].Price) < average.Average(false, cointSell.Index)

	if r1 || r2 || r3 {
		// do sell
		order := &database.Order{
			Contract:  cointSell.Contract,
			Direction: utils.Up,
		}

		if len(orders) > 0 {
			order.SoldPrice = cointSell.Histories[cointSell.Index].Price
			order.SoldTime = cointSell.Histories[cointSell.Index].Time
			err = order.UpdateOrder(ctx)
			if err != nil {
				logrus.Error("update order failed: ", err)
				return err
			}

			if err := order.DeleteOrder(ctx); err != nil {
				logrus.Errorf("Sell order failed, the order detail is %s, the error is %s.", order, err)
				return err
			}
		}
	}

	// for falling
	orders, err = (&database.Order{
		Contract:  cointSell.Contract,
		Direction: utils.Down,
	}).FetchOrder(ctx)
	if err != nil {
		logrus.Error("fetch orders failed: ", err)
		return err
	}

	f1 := average.Average(false, cointSell.Index) >= average.Average(true, cointSell.Index)
	f2 := len(orders) == 1 && utils.StringToFloat64(cointSell.Histories[cointSell.Index].Price) >= utils.StringToFloat64(orders[0].Price)*1.01
	f3 := utils.StringToFloat64(cointSell.Histories[cointSell.Index].Price) > average.Average(false, cointSell.Index)

	if f1 || f2 || f3 {
		// do sell
		order := &database.Order{
			Contract:  cointSell.Contract,
			Direction: utils.Down,
		}

		if len(orders) > 0 {
			order.SoldPrice = cointSell.Histories[cointSell.Index].Price
			order.SoldTime = cointSell.Histories[cointSell.Index].Time
			err = order.UpdateOrder(ctx)
			if err != nil {
				logrus.Error("update order failed: ", err)
				return err
			}

			if err := order.DeleteOrder(ctx); err != nil {
				logrus.Errorf("Sell order failed, the order detail is %s, the error is %s.", order, err)
				return err
			}
		}
	}

	return nil
}
