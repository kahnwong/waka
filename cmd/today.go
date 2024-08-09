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

		// render output
		for _, stat := range stats {
			// print title
			fmt.Println(stat.Title)

			for _, i := range stat.Stats {
				// set slug padding
				var maxStringLength int
				slug := i.Slug
				if len(slug) < maxStringLength {
					slug += strings.Repeat(" ", maxStringLength-len(slug))
				}

				// draw bars
				barArea := 20
				barLength := int(i.Percent) / barArea
				barPadding := barArea - barLength
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
