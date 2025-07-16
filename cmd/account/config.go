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

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get account config info",
	Long:  "Get current account config info.",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.GetConfig()
		client := okx.NewClient(cfg)
		if cfg.OKX.BaseURL != "" {
			client.SetBaseURL(cfg.OKX.BaseURL)
		}
		client.SetSimulated(cfg.OKX.API.IsSimulated)

		info, err := client.GetAccountConfig()
		if err != nil {
			fmt.Fprintf(os.Stderr, "获取账户配置失败: %v\n", err)
			os.Exit(1)
		}

		if viper.GetBool("simple") {
			fmt.Printf("账户ID: %s (母账户ID: %s)\n", info.Uid, info.MainUid)
			fmt.Printf("账户类型: %s, 账户模式: %s, 持仓方式: %s\n", info.Type, info.AcctLv, info.PosMode)
			fmt.Printf("KYC等级: %s, 用户等级: %s\n", info.KycLv, info.Level)
			fmt.Printf("API权限: %s\n", info.Perm)
			fmt.Printf("自动借币: %v, 现货支持借币: %v, 自动还币: %v\n", info.AutoLoan, info.EnableSpotBorrow, info.SpotBorrowAutoRepay)
			fmt.Printf("自成交保护: %s, 逐仓保证金划转: %s\n", info.AcctStpMode, info.CtIsoMode)
			fmt.Printf("期权权限: %s, 希腊字母类型: %s\n", info.OpAuth, info.GreeksType)
			fmt.Printf("API备注: %s, 绑定IP: %s\n", info.Label, info.IP)
			fmt.Printf("角色: %s, 现货角色: %s\n", info.RoleType, info.SpotRoleType)
			fmt.Printf("带单合约: %v\n", info.TraderInsts)
			fmt.Printf("带单币对: %v\n", info.SpotTraderInsts)
		} else {
			output, err := json.MarshalIndent(info, "", "  ")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting output: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(string(output))
		}
	},
}

func init() {
	AccountCmd.AddCommand(configCmd)
}
