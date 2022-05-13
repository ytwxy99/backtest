package cqrs

import (
	"context"
	"fmt"

	"github.com/ytwxy99/autocoins/pkg/configuration"

	"github.com/ytwxy99/backtest/pkg/database"
	"github.com/ytwxy99/backtest/pkg/utils"
)

type Result struct {
	Contract  string
	Direction string
}

func (result *Result) Subscribe(ctx context.Context) error {
	var sum float32
	statistics := map[string]float32{}

	order := &database.Order{
		Contract: result.Contract,
	}

	if order.Contract != utils.ALL {
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

		statistics[order.Contract] = sum

	} else {
		coins, err := utils.ReadLines(ctx.Value("conf").(*configuration.SystemConf).WeightCsv)
		if err != nil {
			return err
		}

		for _, coin := range coins {
			order.Contract = coin
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
			statistics[order.Contract] = sum
		}
	}

	for _, kv := range (&utils.KV{}).MapSortStringFloat(statistics) {
		fmt.Println(kv)
	}

	return nil
}
