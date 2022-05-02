package main

import (
	"context"
	"github.com/sirupsen/logrus"

	"github.com/ytwxy99/backtest/pkg/system"
)

type tt string

func main() {
	ctx := context.Background()
	ctx, err := system.Init(ctx)
	if err != nil {
		logrus.Error("setup mysql failed: ", err)
	}
}
