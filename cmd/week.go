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

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Get summary for week",
	Run: func(cmd *cobra.Command, args []string) {
		if err := wakatime.RenderStats("last_7_days"); err != nil {
			slog.Error("Failed to render stats", "err", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(weekCmd)
}
