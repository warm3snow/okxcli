package asset

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/warm3snow/cexcli/internal/cex"
	"github.com/warm3snow/cexcli/internal/config"
)

var ValuationCmd = &cobra.Command{
	Use:   "valuation [ccy]",
	Short: "Get account asset valuation",
	Long:  `Get account asset valuation. Optionally specify the unit (e.g. BTC, USDT, USD, etc.).\nExample:\n  cexcli valuation\n  cexcli valuation USDT`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		client := cex.NewClient(cfg)

		// Set base URL from config if provided
		if cfg.CEX.BaseURL != "" {
			client.SetBaseURL(cfg.CEX.BaseURL)
		}
		client.SetSimulated(cfg.CEX.API.IsSimulated)

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
