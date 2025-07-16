package market

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

var tickersCmd = &cobra.Command{
	Use:   "tickers",
	Short: "Get all product tickers",
	Long:  "Get all product tickers for a given instType (SPOT, SWAP, FUTURES, OPTION).",
	Run: func(cmd *cobra.Command, args []string) {
		instType, _ := cmd.Flags().GetString("instType")
		uly, _ := cmd.Flags().GetString("uly")
		instFamily, _ := cmd.Flags().GetString("instFamily")

		if instType == "" {
			fmt.Fprintln(os.Stderr, "instType 参数为必填 (SPOT, SWAP, FUTURES, OPTION)")
			os.Exit(1)
		}

		cfg := config.GetConfig()
		client := okx.NewClient(cfg)
		if cfg.OKX.BaseURL != "" {
			client.SetBaseURL(cfg.OKX.BaseURL)
		}
		client.SetSimulated(cfg.OKX.API.IsSimulated)

		tickers, err := client.GetTickers(instType, uly, instFamily)
		if err != nil {
			fmt.Fprintf(os.Stderr, "获取行情失败: %v\n", err)
			os.Exit(1)
		}

		if viper.GetBool("simple") {
			for _, t := range tickers {
				fmt.Printf("产品ID: %s\n", t.InstID)
				fmt.Printf("类型: %s\n", t.InstType)
				fmt.Printf("最新成交价: %s\n", t.Last)
				fmt.Printf("买一价: %s, 卖一价: %s\n", t.BidPx, t.AskPx)
				fmt.Printf("24h最高: %s, 24h最低: %s\n", t.High24h, t.Low24h)
				fmt.Printf("24h成交量(币): %s, 24h成交量(张): %s\n", t.VolCcy24h, t.Vol24h)
				fmt.Printf("时间: %s\n", formatMarketTime(t.Ts))
				fmt.Println()
			}
		} else {
			output, err := json.MarshalIndent(tickers, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "格式化输出失败: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}
	},
}

func formatMarketTime(ts string) string {
	if ts == "" {
		return ""
	}
	t, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return ts
	}
	return time.UnixMilli(t).Format("2006-01-02 15:04:05")
}

func init() {
	tickersCmd.Flags().String("instType", "", "产品类型 (SPOT, SWAP, FUTURES, OPTION) (required)")
	tickersCmd.Flags().String("uly", "", "标的指数 (可选)")
	tickersCmd.Flags().String("instFamily", "", "交易品种 (可选)")
	tickersCmd.MarkFlagRequired("instType")
	MarketCmd.AddCommand(tickersCmd)
}
