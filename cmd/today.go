/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/waka/internal/wakatime"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get summary for today",
	Run: func(cmd *cobra.Command, args []string) {
		if err := wakatime.RenderSummary("Today"); err != nil {
			log.Fatal().Err(err).Msg("Failed to render summary")
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
