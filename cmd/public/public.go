package public

import (
	"github.com/spf13/cobra"
)

var PublicCmd = &cobra.Command{
	Use:   "public",
	Short: "Public data related commands (no authentication required)",
	Long:  `Access public, unauthenticated data such as instruments, tickers, etc.`,
}
