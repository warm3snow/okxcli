package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/warm3snow/cexcli/internal/config"
	"github.com/warm3snow/cexcli/logger"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "cexcli",
		Short: "A CLI tool for interacting with CEX exchange",
		Long: `cexcli is a comprehensive command line interface for interacting with CEX exchange.
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
	cobra.OnInitialize(config.Init)
	cobra.OnInitialize(func() {
		logger.InitLogger(
			config.GetConfig().Logging.Level,
			config.GetConfig().Logging.Format,
		)
	})

	// Global flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./cex.yaml)")

	// Add API key flags (these will override config file values)
	rootCmd.PersistentFlags().String("api-key", "", "CEX API key")
	rootCmd.PersistentFlags().String("api-secret", "", "CEX API secret")
	rootCmd.PersistentFlags().String("passphrase", "", "CEX API passphrase")
	rootCmd.PersistentFlags().Bool("simulated", false, "Use simulated trading")

	// Bind flags to viper
	viper.BindPFlag("cex.api.api_key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("cex.api.secret_key", rootCmd.PersistentFlags().Lookup("api-secret"))
	viper.BindPFlag("cex.api.passphrase", rootCmd.PersistentFlags().Lookup("passphrase"))
	viper.BindPFlag("cex.api.is_simulated", rootCmd.PersistentFlags().Lookup("simulated"))

	// Add commands
	rootCmd.AddCommand(marketCmd)
	rootCmd.AddCommand(balanceCmd)
}
