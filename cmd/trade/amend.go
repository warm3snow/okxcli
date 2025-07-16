package trade

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/warm3snow/okxcli/internal/config"
	"github.com/warm3snow/okxcli/internal/okx"
	"github.com/warm3snow/okxcli/internal/types"
)

var amendCmd = &cobra.Command{
	Use:   "amend",
	Short: "Amend an existing order",
	Long: `Amend an existing order. You must provide either ordId or clOrdId.
For options orders, only one of newPx/newPxUsd/newPxVol can be used and must be consistent with the original order.`,
	Run: func(cmd *cobra.Command, args []string) {
		instId, _ := cmd.Flags().GetString("instId")
		ordId, _ := cmd.Flags().GetString("ordId")
		clOrdId, _ := cmd.Flags().GetString("clOrdId")
		cxlOnFail, _ := cmd.Flags().GetBool("cxlOnFail")
		reqId, _ := cmd.Flags().GetString("reqId")
		newSz, _ := cmd.Flags().GetString("newSz")
		newPx, _ := cmd.Flags().GetString("newPx")
		newPxUsd, _ := cmd.Flags().GetString("newPxUsd")
		newPxVol, _ := cmd.Flags().GetString("newPxVol")
		attachAlgoOrdsStr, _ := cmd.Flags().GetString("attachAlgoOrds")

		if instId == "" {
			fmt.Fprintln(os.Stderr, "instId 为必填参数")
			os.Exit(1)
		}

		if ordId == "" && clOrdId == "" {
			fmt.Fprintln(os.Stderr, "ordId 和 clOrdId 必须至少填写一个")
			os.Exit(1)
		}

		// 检查期权价格参数
		priceParamCount := 0
		if newPx != "" {
			priceParamCount++
		}
		if newPxUsd != "" {
			priceParamCount++
		}
		if newPxVol != "" {
			priceParamCount++
		}
		if priceParamCount > 1 {
			fmt.Fprintln(os.Stderr, "newPx、newPxUsd、newPxVol 只能填写一个")
			os.Exit(1)
		}

		req := &types.AmendOrderRequest{
			InstId:    instId,
			CxlOnFail: cxlOnFail,
			OrdId:     ordId,
			ClOrdId:   clOrdId,
			ReqId:     reqId,
			NewSz:     newSz,
			NewPx:     newPx,
			NewPxUsd:  newPxUsd,
			NewPxVol:  newPxVol,
		}

		// 解析止盈止损参数
		if attachAlgoOrdsStr != "" {
			var attachAlgoOrds []types.AmendOrderAttachAlgo
			if err := json.Unmarshal([]byte(attachAlgoOrdsStr), &attachAlgoOrds); err != nil {
				fmt.Fprintf(os.Stderr, "解析 attachAlgoOrds 参数失败: %v\n", err)
				os.Exit(1)
			}
			req.AttachAlgoOrds = attachAlgoOrds
		}

		cfg := config.GetConfig()
		client := okx.NewClient(cfg)
		if cfg.OKX.BaseURL != "" {
			client.SetBaseURL(cfg.OKX.BaseURL)
		}
		client.SetSimulated(cfg.OKX.API.IsSimulated)

		resp, err := client.AmendOrder(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "修改订单失败: %v\n", err)
			os.Exit(1)
		}

		if viper.GetBool("simple") {
			for _, r := range resp {
				if r.OrdId != "" {
					fmt.Printf("订单ID: %s\n", r.OrdId)
				}
				if r.ClOrdId != "" {
					fmt.Printf("客户订单ID: %s\n", r.ClOrdId)
				}
				if r.ReqId != "" {
					fmt.Printf("修改事件ID: %s\n", r.ReqId)
				}
				fmt.Printf("修改时间: %s\n", r.Ts)
				if r.SMsg != "" {
					fmt.Printf("状态信息: %s\n", r.SMsg)
				}
			}
		} else {
			output, err := json.MarshalIndent(resp, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "格式化输出失败: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}
	},
}

func init() {
	amendCmd.Flags().String("instId", "", "产品ID (必填)")
	amendCmd.Flags().String("ordId", "", "订单ID")
	amendCmd.Flags().String("clOrdId", "", "用户自定义订单ID")
	amendCmd.Flags().Bool("cxlOnFail", false, "修改失败时是否自动撤单")
	amendCmd.Flags().String("reqId", "", "用户自定义修改事件ID")
	amendCmd.Flags().String("newSz", "", "修改的新数量")
	amendCmd.Flags().String("newPx", "", "修改后的新价格")
	amendCmd.Flags().String("newPxUsd", "", "以USD价格进行期权改单")
	amendCmd.Flags().String("newPxVol", "", "以隐含波动率进行期权改单")
	amendCmd.Flags().String("attachAlgoOrds", "", "修改附带止盈止损信息 (JSON格式)")

	TradeCmd.AddCommand(amendCmd)
}
