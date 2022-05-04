package market

import (
	"github.com/ytwxy99/autocoins/pkg/client"
	autils "github.com/ytwxy99/autocoins/pkg/utils"

	"github.com/ytwxy99/backtest/pkg/utils"
)

func InitGateClient() {
	authConf, _ := autils.ReadGateAPIV4(utils.GateConfPath)
	client.GetClient(authConf)
}
