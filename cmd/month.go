/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/waka/internal/wakatime"

	"github.com/spf13/cobra"
)

var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Get summary for month",
	Run: func(cmd *cobra.Command, args []string) {
		wakatime.RenderStats("last_30_days")
	},
}

func init() {
	rootCmd.AddCommand(monthCmd)
}
