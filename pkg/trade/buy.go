package trade

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/ytwxy99/backtest/pkg/database"
)

type BuyOrder struct {
	Contract  string
	Price     string
	Direction string
}

func (buyOrder *BuyOrder) Buy(ctx context.Context) error {
	order := &database.Order{
		Contract:  buyOrder.Contract,
		Price:     buyOrder.Price,
		Direction: buyOrder.Direction,
	}

	orders, err := order.FetchOrder(ctx)
	if err != nil {
		logrus.Errorf("Fetch order failed, the order detail is %s, the error is %s.", order, err)
		return err
	}

	if len(orders) != 0 {
		return nil
	}

	if err := order.AddOrder(ctx); err != nil {
		logrus.Errorf("Add order failed, the order detail is %s, the error is %s.", order, err)
		return err
	}

	return nil
}
