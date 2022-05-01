package main

import (
	"context"
	"github.com/ytwxy99/backtest/pkg/system"
)

func main() {
	ctx := context.Background()
	system.InitCmd(ctx, nil, nil)
}
