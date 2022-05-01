package system

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/ytwxy99/autocoins/pkg/configuration"
	"gorm.io/gorm"
)

// refer: https://github.com/spf13/cobra/blob/v1.2.1/user_guide.md
func InitCmd(ctx context.Context, sysConf *configuration.SystemConf, db *gorm.DB) {
	// init action
	var InitCmd = &cobra.Command{
		Use:   "init [string to echo]",
		Short: "Init back test environment",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	var rootCmd = &cobra.Command{Use: "backtest"}
	rootCmd.AddCommand(InitCmd)
	rootCmd.Execute()
}
