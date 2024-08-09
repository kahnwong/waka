/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/waka/wakatime"
	"github.com/spf13/cobra"
)

var yesterdayCmd = &cobra.Command{
	Use:   "yesterday",
	Short: "Get summary for yesterday",
	Run: func(cmd *cobra.Command, args []string) {
		wakatime.RenderStats("Yesterday")
	},
}

func init() {
	rootCmd.AddCommand(yesterdayCmd)
}
