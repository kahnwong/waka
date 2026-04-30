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

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get summary for today",
	Run: func(cmd *cobra.Command, args []string) {
		if err := wakatime.RenderSummary("Today"); err != nil {
			slog.Error("Failed to render summary", "err", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
