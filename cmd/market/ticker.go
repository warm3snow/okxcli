package market

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/warm3snow/cexcli/internal/cex"
	"github.com/warm3snow/cexcli/internal/config"
)

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

		if cmd.Flag("simple").Value.String() == "true" {
			fmt.Printf("产品ID: %s\n", ticker.InstID)
			fmt.Printf("产品类型: %s\n", ticker.InstType)
			fmt.Printf("最新成交价: %s\n", ticker.Last)
			fmt.Printf("最新成交数量: %s\n", ticker.LastSz)
			fmt.Printf("24小时最高价: %s\n", ticker.High24h)
			fmt.Printf("24小时最低价: %s\n", ticker.Low24h)
			fmt.Printf("24小时开盘价: %s\n", ticker.Open24h)
			fmt.Printf("24小时成交量(币): %s\n", ticker.VolCcy24h)
			fmt.Printf("24小时成交量(张): %s\n", ticker.Vol24h)
			fmt.Printf("时间戳: %s\n", ticker.Ts)
		} else {
			// Pretty print the ticker information
			output, err := json.MarshalIndent(ticker, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
				os.Exit(1)
			}

			fmt.Println(string(output))
		}
	},
}

func init() {
	MarketCmd.AddCommand(tickerCmd)
}
