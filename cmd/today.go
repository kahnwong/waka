/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/waka/internal/wakatime"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get summary for today",
	Run: func(cmd *cobra.Command, args []string) {
		wakatime.RenderSummary("Today")
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
