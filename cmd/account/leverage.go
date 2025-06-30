package account

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

var leverageCmd = &cobra.Command{
	Use:   "leverage",
	Short: "Set leverage for a product",
	Long:  "Set leverage for a product or currency. See OKX API for all scenarios.",
	Run: func(cmd *cobra.Command, args []string) {
		instId, _ := cmd.Flags().GetString("instId")
		ccy, _ := cmd.Flags().GetString("ccy")
		lever, _ := cmd.Flags().GetString("lever")
		mgnMode, _ := cmd.Flags().GetString("mgnMode")
		posSide, _ := cmd.Flags().GetString("posSide")

		if lever == "" || mgnMode == "" {
			fmt.Fprintln(os.Stderr, "lever 和 mgnMode 为必填参数")
			os.Exit(1)
		}
		if instId == "" && ccy == "" {
			fmt.Fprintln(os.Stderr, "instId 和 ccy 至少要传一个")
			os.Exit(1)
		}

		cfg := config.GetConfig()
		client := cex.NewClient(cfg)
		if cfg.CEX.BaseURL != "" {
			client.SetBaseURL(cfg.CEX.BaseURL)
		}
		client.SetSimulated(cfg.CEX.API.IsSimulated)

		req := &types.SetLeverageRequest{
			InstId:  instId,
			Ccy:     ccy,
			Lever:   lever,
			MgnMode: mgnMode,
			PosSide: posSide,
		}
		resp, err := client.SetLeverage(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "设置杠杆失败: %v\n", err)
			os.Exit(1)
		}

		if viper.GetBool("simple") {
			fmt.Printf("产品ID: %s\n", resp.InstId)
			fmt.Printf("杠杆倍数: %s\n", resp.Lever)
			fmt.Printf("保证金模式: %s\n", resp.MgnMode)
			if resp.PosSide != "" {
				fmt.Printf("持仓方向: %s\n", resp.PosSide)
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

func init() {
	leverageCmd.Flags().String("instId", "", "产品ID")
	leverageCmd.Flags().String("ccy", "", "保证金币种")
	leverageCmd.Flags().String("lever", "", "杠杆倍数 (必填)")
	leverageCmd.Flags().String("mgnMode", "", "保证金模式: isolated/cross (必填)")
	leverageCmd.Flags().String("posSide", "", "持仓方向: long/short (仅逐仓开平仓模式必填)")
	AccountCmd.AddCommand(leverageCmd)
}
