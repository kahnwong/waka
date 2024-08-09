/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "Get Wakatime stats for today",
	Run: func(cmd *cobra.Command, args []string) {
		response := getStats("today")
		stats := extractData(response)

		fmt.Println(stats)

		//// prep for output formatting
		//var maxStringLength int
		//for _, i := range category {
		//	if maxStringLength < len(i.Slug) {
		//		maxStringLength = len(i.Slug)
		//	}
		//}
		//
		//// print output
		//fmt.Println("🚀  Projects")
		//for _, i := range category {
		//	slug := i.Slug
		//	if len(slug) < maxStringLength {
		//		slug += strings.Repeat(" ", maxStringLength-len(slug))
		//	}
		//
		//	barArea := 20
		//	barLength := int(i.Percent) / barArea
		//	barPadding := barArea - barLength
		//	bar := fmt.Sprintf("%s%s", strings.Repeat("▇", barLength), strings.Repeat("░", barPadding))
		//	fmt.Printf("%s : %s %v\n", slug, bar, i.Percent)
		//}

	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
}
