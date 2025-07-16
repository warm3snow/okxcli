package trade

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/warm3snow/okxcli/internal/config"
	"github.com/warm3snow/okxcli/internal/okx"
	"github.com/warm3snow/okxcli/internal/types"
)

var cancelOrderCmd = &cobra.Command{
	Use:   "cancel",
	Short: "Cancel an existing order",
	Long:  `Cancel an existing order by providing either ordId or clOrdId along with instId.`,
	Run:   runCancelOrder,
}

var (
	cancelInstId  string
	cancelOrdId   string
	cancelClOrdId string
)

func init() {
	TradeCmd.AddCommand(cancelOrderCmd)

	cancelOrderCmd.Flags().StringVarP(&cancelInstId, "instId", "i", "", "产品ID，如 BTC-USDT")
	cancelOrderCmd.Flags().StringVarP(&cancelOrdId, "ordId", "o", "", "订单ID")
	cancelOrderCmd.Flags().StringVarP(&cancelClOrdId, "clOrdId", "c", "", "用户自定义订单ID")

	cancelOrderCmd.MarkFlagRequired("instId")
}

func runCancelOrder(cmd *cobra.Command, args []string) {
	if cancelOrdId == "" && cancelClOrdId == "" {
		fmt.Fprintln(os.Stderr, "必须提供 ordId 或 clOrdId 中的一个")
		os.Exit(1)
	}

	cfg := config.GetConfig()
	client := okx.NewClient(cfg)
	if cfg.OKX.BaseURL != "" {
		client.SetBaseURL(cfg.OKX.BaseURL)
	}
	client.SetSimulated(cfg.OKX.API.IsSimulated)

	req := &types.CancelOrderRequest{
		InstID:  cancelInstId,
		OrdID:   cancelOrdId,
		ClOrdID: cancelClOrdId,
	}

	results, err := client.CancelOrder(req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "撤单失败: %v\n", err)
		os.Exit(1)
	}

	if cmd.Flag("simple").Value.String() == "true" {
		// 简单模式输出
		for _, result := range results {
			fmt.Printf("订单ID: %s\n", result.OrdID)
			if result.ClOrdID != "" {
				fmt.Printf("自定义订单ID: %s\n", result.ClOrdID)
			}
			fmt.Printf("状态码: %s\n", result.SCode)
			if result.SMsg != "" {
				fmt.Printf("状态信息: %s\n", result.SMsg)
			}
			fmt.Printf("时间戳: %s\n", result.Ts)
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
