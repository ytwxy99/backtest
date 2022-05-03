package driver

import (
	"context"
	"fmt"

	"github.com/ytwxy99/autocoins/pkg/configuration"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/ytwxy99/backtest/pkg/database"
)

func SetupMysql(ctx context.Context) (context.Context, error) {
	conf := ctx.Value("conf").(*configuration.SystemConf).Mysql
	mysqlUrl := fmt.Sprintf("%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.Database)
	db, err := gorm.Open(mysql.Open(mysqlUrl), &gorm.Config{
		//Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return ctx, err
	}

	return context.WithValue(ctx, "DbSession", db), nil
}

func Migrate(ctx context.Context) {
	db := ctx.Value("DbSession").(*gorm.DB)

	db.AutoMigrate(&database.Publish{})
	db.AutoMigrate(&database.HistoryDay{})
}
