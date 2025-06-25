package market

import (
	"github.com/spf13/cobra"
)

var MarketCmd = &cobra.Command{
	Use:   "market",
	Short: "Market data related commands",
	Long: `Market data commands allow you to fetch various market information from CEX exchange,
including tickers, order books, trades, and more.`,
}
