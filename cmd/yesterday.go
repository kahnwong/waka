/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"log/slog"
	"os"

	"github.com/kahnwong/waka/internal/wakatime"
	"github.com/spf13/cobra"
)

var yesterdayCmd = &cobra.Command{
	Use:   "yesterday",
	Short: "Get summary for yesterday",
	Run: func(cmd *cobra.Command, args []string) {
		if err := wakatime.RenderSummary("Yesterday"); err != nil {
			slog.Error("Failed to render summary", "err", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(yesterdayCmd)
}
