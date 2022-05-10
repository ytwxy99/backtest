package cqrs

import (
	"context"
	"fmt"
	"github.com/ytwxy99/backtest/pkg/database"
	"github.com/ytwxy99/backtest/pkg/utils"
)

type Result struct {
	Contract  string
	Direction string
}

func (result *Result) Subscribe(ctx context.Context) error {
	var sum float32
	order := &database.Order{
		Contract: result.Contract,
	}

	orders, err := order.FetchOrderUnscope(ctx)
	if err != nil {
		return err
	}

	for _, order := range orders {
		if order.Direction == utils.Up {
			sum = sum + (utils.StringToFloat32(order.SoldPrice)-utils.StringToFloat32(order.Price))/utils.StringToFloat32(order.Price)
		} else {
			sum = sum + (utils.StringToFloat32(order.Price)-utils.StringToFloat32(order.SoldPrice))/utils.StringToFloat32(order.Price)
		}
	}
	fmt.Println(sum)

	return nil
}
