package trade

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/warm3snow/cexcli/internal/cex"
	"github.com/warm3snow/cexcli/internal/config"
	"github.com/warm3snow/cexcli/internal/types"
)

var orderCmd = &cobra.Command{
	Use:   "order",
	Short: "Place a new order",
	Long:  `Place a new order. Required flags: --instId, --tdMode, --side, --ordType, --sz`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		client := cex.NewClient(cfg)
		if cfg.CEX.BaseURL != "" {
			client.SetBaseURL(cfg.CEX.BaseURL)
		}
		client.SetSimulated(cfg.CEX.API.IsSimulated)

		reduceOnly, _ := cmd.Flags().GetBool("reduceOnly")
		req := &types.PlaceOrderRequest{
			InstID:     mustGetString(cmd, "instId"),
			TdMode:     mustGetString(cmd, "tdMode"),
			Side:       mustGetString(cmd, "side"),
			OrdType:    mustGetString(cmd, "ordType"),
			Sz:         mustGetString(cmd, "sz"),
			Px:         mustGetString(cmd, "px"),
			PosSide:    mustGetString(cmd, "posSide"),
			ClOrdID:    mustGetString(cmd, "clOrdId"),
			ReduceOnly: reduceOnly,
		}

		resp, err := client.PlaceOrder(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error placing order: %v\n", err)
			os.Exit(1)
		}

		if viper.GetBool("simple") {
			for _, order := range resp {
				fmt.Printf("下单成功! 订单ID: %s, 客户自定义ID: %s\n", order.OrdID, order.ClOrdID)
			}
		} else {
			output, err := json.MarshalIndent(resp, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}
	},
}

func mustGetString(cmd *cobra.Command, name string) string {
	val, err := cmd.Flags().GetString(name)
	if err != nil {
		// This should not happen if the flag is defined
		panic(err)
	}
	return val
}

func init() {
	orderCmd.Flags().String("instId", "", "Instrument ID (required)")
	orderCmd.Flags().String("tdMode", "", "Trade mode: cross, isolated, cash, spot_isolated (required)")
	orderCmd.Flags().String("side", "", "Order side: buy, sell (required)")
	orderCmd.Flags().String("ordType", "", "Order type: market, limit, etc. (required)")
	orderCmd.Flags().String("sz", "", "Order size (required)")
	orderCmd.Flags().String("px", "", "Order price (for limit orders)")
	orderCmd.Flags().String("posSide", "", "Position side: long, short, net")
	orderCmd.Flags().String("clOrdId", "", "Client-supplied order ID")
	orderCmd.Flags().Bool("reduceOnly", false, "Reduce-only order")

	orderCmd.MarkFlagRequired("instId")
	orderCmd.MarkFlagRequired("tdMode")
	orderCmd.MarkFlagRequired("side")
	orderCmd.MarkFlagRequired("ordType")
	orderCmd.MarkFlagRequired("sz")

	TradeCmd.AddCommand(orderCmd)
}
