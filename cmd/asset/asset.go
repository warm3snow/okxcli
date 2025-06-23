package asset

import (
	"github.com/spf13/cobra"
)

var AssetCmd = &cobra.Command{
	Use:   "asset",
	Short: "Asset related commands",
}

func init() {
	AssetCmd.AddCommand(BalanceCmd)
	AssetCmd.AddCommand(ValuationCmd)
}
