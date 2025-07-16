package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/warm3snow/okxcli/cmd/account"
	"github.com/warm3snow/okxcli/cmd/asset"
	"github.com/warm3snow/okxcli/cmd/market"
	"github.com/warm3snow/okxcli/cmd/public"
	"github.com/warm3snow/okxcli/cmd/trade"
	"github.com/warm3snow/okxcli/internal/config"
	"github.com/warm3snow/okxcli/logger"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "okxcli",
		Short: "A CLI tool for interacting with OKX exchange",
		Long: `okxcli is a comprehensive command line interface for interacting with OKX exchange.
It provides commands for trading, account management, and market data retrieval.`,
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() {
		config.Init(cfgFile)

		logger.InitLogger(
			config.GetConfig().Logging.Level,
			config.GetConfig().Logging.Format,
		)
	})

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.yaml)")

	// Add API key flags (these will override config file values)
	rootCmd.PersistentFlags().String("api-key", "", "OKX API key")
	rootCmd.PersistentFlags().String("api-secret", "", "OKX API secret")
	rootCmd.PersistentFlags().String("passphrase", "", "OKX API passphrase")
	rootCmd.PersistentFlags().Bool("simulated", false, "Use simulated trading")
	rootCmd.PersistentFlags().Bool("simple", false, "Stdout simple information")

	// Bind flags to viper
	viper.BindPFlag("okx.api.api_key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("okx.api.secret_key", rootCmd.PersistentFlags().Lookup("api-secret"))
	viper.BindPFlag("okx.api.passphrase", rootCmd.PersistentFlags().Lookup("passphrase"))
	viper.BindPFlag("okx.api.is_simulated", rootCmd.PersistentFlags().Lookup("simulated"))
	viper.BindPFlag("simple", rootCmd.PersistentFlags().Lookup("simple"))

	// Add commands
	rootCmd.AddCommand(market.MarketCmd)
	rootCmd.AddCommand(asset.AssetCmd)
	rootCmd.AddCommand(account.AccountCmd)
	rootCmd.AddCommand(trade.TradeCmd)
	rootCmd.AddCommand(public.PublicCmd)
}
