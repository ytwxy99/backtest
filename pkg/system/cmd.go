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
			if len(args) != 0 {
				if len(args) > 1 {
					fmt.Println("only one coin can be entered!")
					return
				}
				(&cqrs.Result{
					Contract:  args[0],
					Direction: utils.Up,
				}).Subscribe(ctx)
			} else {
				(&cqrs.Result{
					Contract:  "BTC_USDT",
					Direction: utils.Up,
				}).Subscribe(ctx)
			}
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
	var Coint4hCmd = &cobra.Command{
		Use:   "4h [string to echo]",
		Short: "Using coint 4h policy to do a test",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				if len(args) > 1 {
					fmt.Println("only one coin can be entered!")
					return
				}
				(&cqrs.PublishBus{
					Contract: args[0],
					Event:    "4h",
					Status:   utils.NewPublish,
				}).Publish(ctx)
			} else {
				(&cqrs.PublishBus{
					Contract: "BTC_USDT",
					Event:    "4h",
					Status:   utils.NewPublish,
				}).Publish(ctx)
			}

		},
	}

	// use trend policy
	var Coint30mCmd = &cobra.Command{
		Use:   "30m [string to echo]",
		Short: "Using coint 30m policy to do a test",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				if len(args) > 1 {
					fmt.Println("only one coin can be entered!")
					return
				}
				(&cqrs.PublishBus{
					Contract: args[0],
					Event:    "30m",
					Status:   utils.NewPublish,
				}).Publish(ctx)
			} else {
				(&cqrs.PublishBus{
					Contract: "BTC_USDT",
					Event:    "30m",
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
	TestCmd.AddCommand(Coint4hCmd)
	TestCmd.AddCommand(Coint30mCmd)
	rootCmd.Execute()
}
