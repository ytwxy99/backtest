package system

import (
	"context"
	"fmt"

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
			(&cqrs.PublishBus{
				Contract: "*",
				Event:    "init",
				Status:   utils.NewPublish,
			}).Publish(ctx)
		},
	}

	var Result = &cobra.Command{
		Use:   "result [string to echo]",
		Short: "back test result",
		Run: func(cmd *cobra.Command, args []string) {
			(&cqrs.Result{
				Contract:  "BTC_USDT",
				Direction: utils.Up,
			}).Subscribe(ctx)
		},
	}

	var Subscribe = &cobra.Command{
		Use:   "subscribe [string to echo]",
		Short: "Start Subscribe",
		Run: func(cmd *cobra.Command, args []string) {
			(&cqrs.SubscribeBus{}).Subscribe(ctx)
		},
	}

	var TestCmd = &cobra.Command{
		Use:   "test [string to echo]",
		Short: "Do a test which you can choose",
		Args:  cobra.MinimumNArgs(1),
	}

	// use trend policy
	var CointegrationCmd = &cobra.Command{
		Use:   "coint [string to echo]",
		Short: "Using cointegration policy to do a test",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				if len(args) > 1 {
					fmt.Println("only one coin can be entered!")
					return
				}
				(&cqrs.PublishBus{
					Contract: args[0],
					Event:    "cointegration",
					Status:   utils.NewPublish,
				}).Publish(ctx)
			} else {
				(&cqrs.PublishBus{
					Contract: "BTC_USDT",
					Event:    "cointegration",
					Status:   utils.NewPublish,
				}).Publish(ctx)
			}

		},
	}

	var rootCmd = &cobra.Command{Use: "backtest"}
	rootCmd.AddCommand(InitCmd)
	rootCmd.AddCommand(Result)
	rootCmd.AddCommand(Subscribe)
	rootCmd.AddCommand(TestCmd)
	TestCmd.AddCommand(CointegrationCmd)
	rootCmd.Execute()
}
