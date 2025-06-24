package trade

import (
	"github.com/spf13/cobra"
)

var TradeCmd = &cobra.Command{
	Use:   "trade",
	Short: "Trade related commands (e.g. place order)",
	Long:  `Provides commands for trading operations like placing, canceling, or amending orders.`,
}

func init() {
	// 子命令将在这里注册
}
