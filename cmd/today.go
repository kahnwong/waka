/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get Wakatime stats for today",
	Run: func(cmd *cobra.Command, args []string) {
		response := getStats("today")
		stats := extractData(response)

		// get info for slug padding
		var maxStringLength int
		for _, stat := range stats {
			for _, i := range stat.Stats {
				// set slug padding
				slug := i.Slug
				if maxStringLength < len(slug) {
					maxStringLength = len(slug)
				}
			}
		}

		//render output
		for _, stat := range stats {
			// print title
			fmt.Println(stat.Title)

			for _, i := range stat.Stats {
				// set slug padding
				slug := i.Slug
				if len(slug) < maxStringLength {
					slug += strings.Repeat(" ", maxStringLength-len(slug))
				}

				// draw bars
				barLength := (int(i.Percent) / 10) * 2
				barPadding := (20 - barLength)
				bar := fmt.Sprintf("%s%s", strings.Repeat("▇", barLength), strings.Repeat("░", barPadding))
				fmt.Printf("%s : %s %v\n", slug, bar, i.Percent)
			}

			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
