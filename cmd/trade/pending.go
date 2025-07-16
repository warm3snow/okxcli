package trade

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/warm3snow/okxcli/internal/config"
	"github.com/warm3snow/okxcli/internal/okx"
)

var pendingOrdersCmd = &cobra.Command{
	Use:   "pending",
	Short: "Get pending orders",
	Long:  `Get a list of all pending orders. You can filter by instType, instId, ordType, etc.`,
	Run:   runPendingOrders,
}

var (
	pendingInstType string
	pendingInstId   string
	pendingOrdType  string
	pendingState    string
	pendingAfter    string
	pendingBefore   string
	pendingLimit    string
)

func init() {
	TradeCmd.AddCommand(pendingOrdersCmd)

	pendingOrdersCmd.Flags().StringVarP(&pendingInstType, "instType", "t", "", "产品类型，如 SPOT, MARGIN, SWAP, FUTURES, OPTION")
	pendingOrdersCmd.Flags().StringVarP(&pendingInstId, "instId", "i", "", "产品ID，如 BTC-USDT")
	pendingOrdersCmd.Flags().StringVarP(&pendingOrdType, "ordType", "o", "", "订单类型，如 market, limit, post_only, fok, ioc")
	pendingOrdersCmd.Flags().StringVarP(&pendingState, "state", "s", "", "订单状态，如 live, partially_filled")
	pendingOrdersCmd.Flags().StringVar(&pendingAfter, "after", "", "请求此ID之前（更旧的数据）的分页内容")
	pendingOrdersCmd.Flags().StringVar(&pendingBefore, "before", "", "请求此ID之后（更新的数据）的分页内容")
	pendingOrdersCmd.Flags().StringVarP(&pendingLimit, "limit", "l", "", "返回结果的数量，最大为100条")
}

func runPendingOrders(cmd *cobra.Command, args []string) {
	cfg := config.GetConfig()
	client := okx.NewClient(cfg)
	if cfg.OKX.BaseURL != "" {
		client.SetBaseURL(cfg.OKX.BaseURL)
	}
	client.SetSimulated(cfg.OKX.API.IsSimulated)

	results, err := client.GetPendingOrders(
		pendingInstType,
		pendingInstId,
		pendingOrdType,
		pendingState,
		pendingAfter,
		pendingBefore,
		pendingLimit,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "获取未完成订单失败: %v\n", err)
		os.Exit(1)
	}

	if cmd.Flag("simple").Value.String() == "true" {
		// 简单模式输出
		for _, order := range results {
			fmt.Printf("订单ID: %s\n", order.OrdId)
			fmt.Printf("产品: %s (%s)\n", order.InstId, order.InstType)
			fmt.Printf("类型: %s\n", order.OrdType)
			fmt.Printf("方向: %s\n", order.Side)
			fmt.Printf("价格: %s\n", order.Px)
			fmt.Printf("数量: %s\n", order.Sz)
			fmt.Printf("已成交: %s\n", order.AccFillSz)
			fmt.Printf("状态: %s\n", order.State)
			fmt.Printf("创建时间: %s\n", formatTime(order.CTime))
			fmt.Printf("更新时间: %s\n", formatTime(order.UTime))
			if order.Fee != "0" {
				fmt.Printf("手续费: %s %s\n", order.Fee, order.FeeCcy)
			}
			if order.Rebate != "0" {
				fmt.Printf("返佣: %s %s\n", order.Rebate, order.RebateCcy)
			}
			fmt.Println()
		}
	} else {
		// 完整JSON输出
		output, err := json.MarshalIndent(results, "", "  ")
		if err != nil {
			fmt.Fprintf(os.Stderr, "格式化输出失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(string(output))
	}
}

func formatTime(timestamp string) string {
	if timestamp == "" {
		return ""
	}
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return timestamp
	}
	return time.UnixMilli(ts).Format("2006-01-02 15:04:05.000")
}
