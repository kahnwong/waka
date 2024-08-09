/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/carlmjohnson/requests"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get Wakatime stats for today",
	Run: func(cmd *cobra.Command, args []string) {
		var response WakatimeStats
		err := requests.
			URL(WakatimeEndpoint).
			Method(http.MethodGet).
			Path("api/v1/users/current/summaries").
			Param("range", "today").
			Header("Authorization", createAuthorizationHeader()).
			ToJSON(&response).
			Fetch(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Failed to get stats for today")
		}

		// init data structs
		var category []Stats

		// extract values
		for _, i := range response.Data {
			for _, l := range i.Languages {
				category = append(category, Stats{
					Slug:    fmt.Sprintf("%s (%s)", l.Name, l.Text),
					Percent: l.Percent,
				})
			}
		}

		// prep for output formatting
		var maxStringLength int
		for _, i := range category {
			if maxStringLength < len(i.Slug) {
				maxStringLength = len(i.Slug)
			}
		}

		// print output
		fmt.Println("🚀  Projects")
		for _, i := range category {
			slug := i.Slug
			if len(slug) < maxStringLength {
				slug += strings.Repeat(" ", maxStringLength-len(slug))
			}

			barArea := 20
			barLength := int(i.Percent) / barArea
			barPadding := barArea - barLength
			bar := fmt.Sprintf("%s%s", strings.Repeat("▇", barLength), strings.Repeat("░", barPadding))
			fmt.Printf("%s : %s %v\n", slug, bar, i.Percent)
		}

		// project
		// language
		// editors
		// os
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
