/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"github.com/kahnwong/waka/internal/wakatime"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var yesterdayCmd = &cobra.Command{
	Use:   "yesterday",
	Short: "Get summary for yesterday",
	Run: func(cmd *cobra.Command, args []string) {
		if err := wakatime.RenderSummary("Yesterday"); err != nil {
			log.Fatal().Err(err).Msg("Failed to render summary")
		}
	},
}

func init() {
	rootCmd.AddCommand(yesterdayCmd)
}
