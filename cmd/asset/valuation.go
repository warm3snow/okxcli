package asset

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/warm3snow/okxcli/internal/config"
	"github.com/warm3snow/okxcli/internal/okx"
)

var ValuationCmd = &cobra.Command{
	Use:   "valuation [ccy]",
	Short: "Get account asset valuation",
	Long:  `Get account asset valuation. Optionally specify the unit (e.g. BTC, USDT, USD, etc.).\nExample:\n  okxcli valuation\n  okxcli valuation USDT`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		client := okx.NewClient(cfg)

		// Set base URL from config if provided
		if cfg.OKX.BaseURL != "" {
			client.SetBaseURL(cfg.OKX.BaseURL)
		}
		client.SetSimulated(cfg.OKX.API.IsSimulated)

		ccy := ""
		if len(args) > 0 {
			ccy = args[0]
		}

		valuation, err := client.GetAssetValuation(ccy)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching asset valuation: %v\n", err)
			os.Exit(1)
		}

		output, err := json.MarshalIndent(valuation, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
	},
}
