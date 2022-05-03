package system

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/ytwxy99/autocoins/pkg/utils"

	"github.com/ytwxy99/backtest/pkg/database/driver"
)

func Init(ctx context.Context) (context.Context, error) {
	sysConf, err := utils.ReadSystemConfig("./etc/autoCoin.yml")
	if err != nil {
		logrus.Error("read configure file failed, ", err)
	}

	utils.InitLog(sysConf.LogPath)
	ctx = context.WithValue(ctx, "conf", sysConf)
	ctx, err = driver.SetupMysql(ctx)
	if err != nil {
		return ctx, err
	}

	driver.Migrate(ctx)

	InitCmd(ctx)

	return ctx, nil
}