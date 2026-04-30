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

var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Get summary for month",
	Run: func(cmd *cobra.Command, args []string) {
		if err := wakatime.RenderStats("last_30_days"); err != nil {
			slog.Error("Failed to render stats", "err", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(monthCmd)
}
