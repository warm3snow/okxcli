package account

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/warm3snow/cexcli/internal/cex"
	"github.com/warm3snow/cexcli/internal/config"
)

var PositionsCmd = &cobra.Command{
	Use:   "positions",
	Short: "Get account positions",
	Long:  "Get account positions. 可用参数: --instId=BTC-USDT --instType=SWAP ...",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		client := cex.NewClient(cfg)
		if cfg.CEX.BaseURL != "" {
			client.SetBaseURL(cfg.CEX.BaseURL)
		}
		client.SetSimulated(cfg.CEX.API.IsSimulated)

		params := map[string]string{}
		instId, _ := cmd.Flags().GetString("instId")
		instType, _ := cmd.Flags().GetString("instType")
		posId, _ := cmd.Flags().GetString("posId")
		if instId != "" {
			params["instId"] = instId
		}
		if instType != "" {
			params["instType"] = instType
		}
		if posId != "" {
			params["posId"] = posId
		}

		positions, err := client.GetAccountPositions(params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching positions: %v\n", err)
			os.Exit(1)
		}
		if viper.GetBool("simple") {
			for _, p := range positions {
				fmt.Printf("产品: %s, 持仓方向: %s, 持仓数量: %s, 杠杆: %s, 未实现盈亏: %s, 保证金模式: %s\n",
					p.InstId, p.PosSide, p.Pos, p.Lever, p.Upl, p.MgnMode)
			}
		} else {
			output, err := json.MarshalIndent(positions, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}
	},
}

func init() {
	PositionsCmd.Flags().String("instId", "", "产品ID，如BTC-USDT")
	PositionsCmd.Flags().String("instType", "", "产品类型，如SWAP")
	PositionsCmd.Flags().String("posId", "", "持仓ID")
}
