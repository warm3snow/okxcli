package account

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/warm3snow/okxcli/internal/config"
	"github.com/warm3snow/okxcli/internal/okx"
)

var PositionRiskCmd = &cobra.Command{
	Use:   "position-risk",
	Short: "Get account position risk",
	Long:  "Get account position risk. 可用参数: --instType=MARGIN/SWAP/FUTURES/OPTION",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		client := okx.NewClient(cfg)
		if cfg.OKX.BaseURL != "" {
			client.SetBaseURL(cfg.OKX.BaseURL)
		}
		client.SetSimulated(cfg.OKX.API.IsSimulated)

		instType, _ := cmd.Flags().GetString("instType")
		risk, err := client.GetAccountPositionRisk(instType)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error fetching position risk: %v\n", err)
			os.Exit(1)
		}
		if viper.GetBool("simple") {
			for _, r := range risk {
				fmt.Printf("总有效保证金: %s\n", r.AdjEq)
				for _, b := range r.BalData {
					fmt.Printf("币种: %s, 总权益: %s, 美金权益: %s\n", b.Ccy, b.Eq, b.DisEq)
				}
				for _, p := range r.PosData {
					fmt.Printf("产品: %s, 持仓方向: %s, 持仓数量: %s, 保证金模式: %s\n",
						p.InstId, p.PosSide, p.Pos, p.MgnMode)
				}
			}
		} else {
			output, err := json.MarshalIndent(risk, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}
	},
}

func init() {
	PositionRiskCmd.Flags().String("instType", "", "产品类型 MARGIN/SWAP/FUTURES/OPTION")
}
