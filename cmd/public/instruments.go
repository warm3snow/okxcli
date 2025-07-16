package public

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/warm3snow/okxcli/internal/config"
	"github.com/warm3snow/okxcli/internal/okx"
)

var instrumentsCmd = &cobra.Command{
	Use:   "instruments",
	Short: "Get instrument information",
	Long:  "Get a list of all tradable instruments. Requires --instType.",
	Run: func(cmd *cobra.Command, args []string) {
		instType, _ := cmd.Flags().GetString("instType")
		uly, _ := cmd.Flags().GetString("uly")
		instFamily, _ := cmd.Flags().GetString("instFamily")
		instId, _ := cmd.Flags().GetString("instId")

		if instType == "" {
			fmt.Fprintln(os.Stderr, "instType is a required parameter (e.g., SPOT, SWAP, FUTURES, OPTION)")
			os.Exit(1)
		}

		cfg := config.GetConfig()
		client := okx.NewClient(cfg)
		if cfg.OKX.BaseURL != "" {
			client.SetBaseURL(cfg.OKX.BaseURL)
		}
		client.SetSimulated(cfg.OKX.API.IsSimulated)

		instruments, err := client.GetInstruments(instType, uly, instFamily, instId)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching instruments: %v\n", err)
			os.Exit(1)
		}

		if viper.GetBool("simple") {
			for _, i := range instruments {
				fmt.Printf("产品ID: %s, 状态: %s\n", i.InstID, i.State)
				if i.BaseCcy != "" {
					fmt.Printf("  交易货币: %s, 计价货币: %s\n", i.BaseCcy, i.QuoteCcy)
				}
				if i.SettleCcy != "" {
					fmt.Printf("  结算币种: %s, 合约面值: %s %s\n", i.SettleCcy, i.CtVal, i.CtValCcy)
				}
				fmt.Printf("  下单精度: (价格: %s, 数量: %s), 最小下单量: %s\n", i.TickSz, i.LotSz, i.MinSz)
				if i.Lever != "" {
					fmt.Printf("  最大杠杆: %s\n", i.Lever)
				}
				listTime, _ := strconv.ParseInt(i.ListTime, 10, 64)
				fmt.Printf("  上线时间: %s\n\n", time.UnixMilli(listTime).Format("2006-01-02"))
			}
		} else {
			output, err := json.MarshalIndent(instruments, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}
	},
}

func init() {
	instrumentsCmd.Flags().String("instType", "", "Product type (e.g., SPOT, SWAP) (required)")
	instrumentsCmd.Flags().String("uly", "", "Underlying index (for FUTURES/SWAP/OPTION)")
	instrumentsCmd.Flags().String("instFamily", "", "Instrument family (for FUTURES/SWAP/OPTION)")
	instrumentsCmd.Flags().String("instId", "", "Instrument ID")
	// instrumentsCmd.MarkFlagRequired("instType")
	PublicCmd.AddCommand(instrumentsCmd)
}
