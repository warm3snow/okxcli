package account

import "github.com/spf13/cobra"

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account related commands",
}

func init() {
	AccountCmd.AddCommand(BalanceCmd)
	AccountCmd.AddCommand(PositionsCmd)
	AccountCmd.AddCommand(PositionRiskCmd)
}
