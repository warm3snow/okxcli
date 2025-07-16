package account

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

var positionModeCmd = &cobra.Command{
	Use:   "position-mode",
	Short: "Set position mode",
	Long:  "Set position mode for futures and perpetual contracts. long_short_mode: long/short positions allowed, net_mode: single direction only",
	Run: func(cmd *cobra.Command, args []string) {
		posMode, _ := cmd.Flags().GetString("posMode")

		if posMode == "" {
			fmt.Fprintln(os.Stderr, "posMode 为必填参数")
			os.Exit(1)
		}

		if posMode != "long_short_mode" && posMode != "net_mode" {
			fmt.Fprintln(os.Stderr, "posMode 必须为 long_short_mode 或 net_mode")
			os.Exit(1)
		}

		cfg := config.GetConfig()
		client := okx.NewClient(cfg)
		if cfg.OKX.BaseURL != "" {
			client.SetBaseURL(cfg.OKX.BaseURL)
		}
		client.SetSimulated(cfg.OKX.API.IsSimulated)

		req := &types.SetPositionModeRequest{
			PosMode: posMode,
		}
		resp, err := client.SetPositionMode(req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "设置持仓模式失败: %v\n", err)
			os.Exit(1)
		}

		if viper.GetBool("simple") {
			fmt.Printf("持仓模式: %s\n", resp.PosMode)
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
	positionModeCmd.Flags().String("posMode", "", "持仓模式: long_short_mode(开平仓模式)/net_mode(买卖模式) (必填)")
	AccountCmd.AddCommand(positionModeCmd)
}
