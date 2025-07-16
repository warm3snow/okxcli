package asset

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/warm3snow/okxcli/internal/config"
	"github.com/warm3snow/okxcli/internal/okx"
)

var BalanceCmd = &cobra.Command{
	Use:   "balance [currencies...]",
	Short: "Get account balance information",
	Long: `Get account balance information for all currencies or specified currencies.
Example:
  okxcli balance            # Get all balances
  okxcli balance BTC,USDT  # Get balances for specific currencies`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		client := okx.NewClient(cfg)

		// Set base URL from config if provided
		if cfg.OKX.BaseURL != "" {
			client.SetBaseURL(cfg.OKX.BaseURL)
		}

		// Set simulated trading mode
		client.SetSimulated(cfg.OKX.API.IsSimulated)

		// Parse currencies from args
		var currencies []string
		if len(args) > 0 {
			currencies = strings.Split(args[0], ",")
		}

		balances, err := client.GetBalances(currencies...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching balances: %v\n", err)
			os.Exit(1)
		}

		// Pretty print the balance information
		output, err := json.MarshalIndent(balances, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(output))
	},
}
