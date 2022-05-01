package system

import (
	"github.com/ytwxy99/autocoins/database"
	"github.com/ytwxy99/autocoins/pkg/client"
	"github.com/ytwxy99/autocoins/pkg/configuration"
	"github.com/ytwxy99/autocoins/pkg/utils"
)

// init system base data
func Init(authConf *configuration.GateAPIV4, sysConf *configuration.SystemConf) {
	_, ctx := client.GetClient(authConf)
	utils.InitLog(sysConf.LogPath)
	db := database.GetDB(sysConf)
	database.InitDB(db)
	InitCmd(ctx, sysConf, db)
}
