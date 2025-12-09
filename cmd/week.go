/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/waka/internal/wakatime"
	"github.com/spf13/cobra"
)

var weekCmd = &cobra.Command{
	Use:   "week",
	Short: "Get summary for week",
	Run: func(cmd *cobra.Command, args []string) {
		wakatime.RenderStats("last_7_days")
	},
}

func init() {
	rootCmd.AddCommand(weekCmd)
}
