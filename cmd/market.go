package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/warm3snow/cexcli/internal/cex"
	"github.com/warm3snow/cexcli/internal/config"
)

var marketCmd = &cobra.Command{
	Use:   "market",
	Short: "Market data related commands",
	Long: `Market data commands allow you to fetch various market information from CEX exchange,
including tickers, order books, trades, and more.`,
}

var tickerCmd = &cobra.Command{
	Use:   "ticker [symbol]",
	Short: "Get ticker information for a trading pair",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		symbol := args[0]

		cfg := config.GetConfig()
		client := cex.NewClient(cfg)

		// Set base URL from config if provided
		if cfg.CEX.BaseURL != "" {
			client.SetBaseURL(cfg.CEX.BaseURL)
		}

		// Set simulated trading mode
		client.SetSimulated(cfg.CEX.API.IsSimulated)

		ticker, err := client.GetTicker(symbol)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching ticker: %v\n", err)
			os.Exit(1)
		}

		// Pretty print the ticker information
		output, err := json.MarshalIndent(ticker, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(marketCmd)
	marketCmd.AddCommand(tickerCmd)
}
