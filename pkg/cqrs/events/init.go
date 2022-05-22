package events

import (
	"context"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/ytwxy99/autocoins/pkg/configuration"
	"github.com/ytwxy99/autocoins/pkg/interfaces"
	autils "github.com/ytwxy99/autocoins/pkg/utils"

	"github.com/ytwxy99/backtest/pkg/database"
	"github.com/ytwxy99/backtest/pkg/utils"
	"github.com/ytwxy99/backtest/pkg/utils/market"
)

type InitEvent struct {
	EventMetadata map[string]string
}

// DoEvent initialize back test system
func (initEvent *InitEvent) DoEvent(ctx context.Context) (string, error) {
	var writeLines []string
	futureSort := map[string]int64{}
	conf := ctx.Value("conf").(*configuration.SystemConf)
	market.InitGateClient()

	futures, err := (&interfaces.Future{
		Settle: utils.Settle,
	}).GetAllFutures(ctx)
	if err != nil {
		logrus.Error("get all futures failed: ", err)
	}

	for _, future := range futures {
		futureSort[future.Name] = future.TradeSize
	}

	for index, future := range (&utils.KVs2i64{}).MapSortStringInt64(futureSort) {
		if index < 30 {
			writeLines = append(writeLines, future.Key)
		}
	}
	autils.WriteLines(writeLines, conf.WeightCsv)

	coins, err := autils.ReadLines(conf.WeightCsv)
	if err != nil {
		logrus.Error("read file failed: ", err)
	}

	// 4h market
	for _, coin := range coins {
		markets := (&interfaces.MarketArgs{
			CurrencyPair: coin,
			Interval:     utils.Intervel4H,
			Level:        utils.Level4Hour,
		}).SpotMarket()

		for _, market := range markets {
			timeTrans, err := time.ParseInLocation("2006-01-02 15:04:00", market[0], time.Local)
			if err != nil {
				logrus.Error("time trans failed: ", err)
				return err.Error(), err
			}

			history := &database.HistoryFourHour{
				Contract: coin,
				Price:    market[2],
				Time:     timeTrans,
			}
			if err = history.AddHistoryFourHour(ctx); err != nil {
				if !strings.Contains(err.Error(), utils.DBHistoryUniq) {
					logrus.Error("add 4h history record failed: ", err)
					return err.Error(), err
				}
			}
		}
	}

	// 30m market
	for _, coin := range coins {
		markets := (&interfaces.MarketArgs{
			CurrencyPair: coin,
			Interval:     utils.Intervel30M,
			Level:        utils.Level30Min,
		}).SpotMarket()

		for _, market := range markets {
			timeTrans, err := time.ParseInLocation("2006-01-02 15:04:00", market[0], time.Local)
			if err != nil {
				logrus.Error("time trans failed: ", err)
				return err.Error(), err
			}

			history := &database.HistoryThirtyMin{
				Contract: coin,
				Price:    market[2],
				Time:     timeTrans,
			}
			if err = history.AddHistoryThirtyMin(ctx); err != nil {
				if !strings.Contains(err.Error(), utils.DBHistoryUniq) {
					logrus.Error("add 4h history record failed: ", err)
					return err.Error(), err
				}
			}
		}
	}

	return "initialize back test success", nil
}
