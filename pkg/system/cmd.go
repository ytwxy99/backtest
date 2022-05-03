package system

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/ytwxy99/backtest/pkg/cqrs"
	"github.com/ytwxy99/backtest/pkg/utils"
)

// refer: https://github.com/spf13/cobra/blob/v1.2.1/user_guide.md
func InitCmd(ctx context.Context) {
	var InitCmd = &cobra.Command{
		Use:   "init [string to echo]",
		Short: "Init back test environment",
		Run: func(cmd *cobra.Command, args []string) {
			//todo(wangxiaoyu), test case
			(&cqrs.PublishBus{
				Contract: "BTC_USDT",
				Event:    "cointegration",
				Status:   utils.NewPublish,
			}).Publish(ctx)
		},
	}

	var Subscribe = &cobra.Command{
		Use:   "subscribe [string to echo]",
		Short: "Start Subscribe",
		Run: func(cmd *cobra.Command, args []string) {
			(&cqrs.SubscribeBus{}).Subscribe(ctx)
		},
	}

	var rootCmd = &cobra.Command{Use: "backtest"}
	rootCmd.AddCommand(InitCmd)
	rootCmd.AddCommand(Subscribe)
	rootCmd.Execute()
}
