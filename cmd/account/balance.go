package account

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/warm3snow/cexcli/internal/cex"
	"github.com/warm3snow/cexcli/internal/config"
)

var BalanceCmd = &cobra.Command{
	Use:   "balance [ccy1,ccy2,...]",
	Short: "Get trading account balance",
	Long:  "Get trading account balance. Optionally specify currencies, e.g. cexcli asset account BTC,ETH",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		client := cex.NewClient(cfg)
		if cfg.CEX.BaseURL != "" {
			client.SetBaseURL(cfg.CEX.BaseURL)
		}
		client.SetSimulated(cfg.CEX.API.IsSimulated)

		var ccys []string
		if len(args) > 0 && args[0] != "" {
			ccys = strings.Split(args[0], ",")
		}
		balances, err := client.GetAccountBalance(ccys...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching account balance: %v\n", err)
			os.Exit(1)
		}
		output, err := json.MarshalIndent(balances, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
	},
}
