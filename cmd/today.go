/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get summary for today",
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
			color.Green(stat.Title)

			for _, i := range stat.Stats {
				// set slug padding
				slug := i.Slug
				if len(slug) < maxStringLength {
					slug += strings.Repeat(" ", maxStringLength-len(slug))
				}

				// draw bars
				barLength := int(i.Percent) / 5 // divided by 5 to reduce rendered bar characters
				barPadding := 20 - barLength    // deduct from 20 as it would be full bar length based on x/5 from barLength
				bar := fmt.Sprintf("%s%s", strings.Repeat("▇", barLength), strings.Repeat("░", barPadding))

				blue := color.New(color.FgBlue).SprintFunc()
				fmt.Printf("%s : %s %v%%\n", blue(slug), bar, i.Percent)
			}

			fmt.Println("")
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
