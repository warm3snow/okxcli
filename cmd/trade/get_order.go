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
	"github.com/warm3snow/okxcli/internal/types"
)

var getOrderCmd = &cobra.Command{
	Use:   "get-order",
	Short: "Get order details",
	Long:  `Get details of a single order. Requires instId and either ordId or clOrdId.`,
	Run:   runGetOrder,
}

var (
	getOrderInstId  string
	getOrderOrdId   string
	getOrderClOrdId string
)

func init() {
	TradeCmd.AddCommand(getOrderCmd)

	getOrderCmd.Flags().StringVarP(&getOrderInstId, "instId", "i", "", "产品ID，如 BTC-USDT (required)")
	getOrderCmd.Flags().StringVarP(&getOrderOrdId, "ordId", "o", "", "订单ID")
	getOrderCmd.Flags().StringVarP(&getOrderClOrdId, "clOrdId", "c", "", "用户自定义订单ID")

	getOrderCmd.MarkFlagRequired("instId")
}

func runGetOrder(cmd *cobra.Command, args []string) {
	if getOrderOrdId == "" && getOrderClOrdId == "" {
		fmt.Fprintln(os.Stderr, "必须提供 ordId 或 clOrdId 中的一个")
		os.Exit(1)
	}

	cfg := config.GetConfig()
	client := okx.NewClient(cfg)
	if cfg.OKX.BaseURL != "" {
		client.SetBaseURL(cfg.OKX.BaseURL)
	}
	client.SetSimulated(cfg.OKX.API.IsSimulated)

	results, err := client.GetOrder(
		getOrderInstId,
		getOrderOrdId,
		getOrderClOrdId,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "获取订单信息失败: %v\n", err)
		os.Exit(1)
	}

	if cmd.Flag("simple").Value.String() == "true" {
		// 简单模式输出
		for _, order := range results {
			printOrderSimple(&order)
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

func printOrderSimple(order *types.PendingOrdersResponse) {
	fmt.Printf("订单ID: %s\n", order.OrdId)
	if order.ClOrdId != "" {
		fmt.Printf("自定义ID: %s\n", order.ClOrdId)
	}
	fmt.Printf("产品: %s (%s)\n", order.InstId, order.InstType)
	fmt.Printf("订单类型: %s\n", order.OrdType)
	fmt.Printf("方向: %s (%s)\n", order.Side, order.PosSide)
	fmt.Printf("委托价格: %s\n", order.Px)
	fmt.Printf("委托数量: %s\n", order.Sz)
	fmt.Printf("成交均价: %s\n", order.AvgPx)
	fmt.Printf("累计成交: %s\n", order.AccFillSz)
	fmt.Printf("状态: %s\n", order.State)
	fmt.Printf("创建时间: %s\n", formatOrderTime(order.CTime))
	fmt.Printf("更新时间: %s\n", formatOrderTime(order.UTime))
	if order.Fee != "0" {
		fmt.Printf("手续费: %s %s\n", order.Fee, order.FeeCcy)
	}
	if order.Rebate != "0" {
		fmt.Printf("返佣: %s %s\n", order.Rebate, order.RebateCcy)
	}
	fmt.Println()
}

func formatOrderTime(timestamp string) string {
	if timestamp == "" {
		return "N/A"
	}
	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return timestamp
	}
	return time.UnixMilli(ts).Format("2006-01-02 15:04:05")
}
