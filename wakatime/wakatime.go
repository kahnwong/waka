/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package wakatime

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/viper"
)

var apiEndpoint = "https://wakatime.com"

// api response
type categoryStats []struct {
	Name    string  `json:"name"`
	Text    string  `json:"text"`
	Percent float64 `json:"percent"`
}

// parsed structs
type parsedCategoryStats struct {
	Slug    string
	Percent float64
}

type parsedStats struct {
	Title string
	Stats []parsedCategoryStats
}

func createAuthorizationHeader() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(viper.GetString("WAKATIME_API_KEY"))))
}

func appendToKey(category string, keyStats categoryStats) parsedStats {
	var keyStatsParsed []parsedCategoryStats

	for _, i := range keyStats {
		keyStatsParsed = append(keyStatsParsed, parsedCategoryStats{
			Slug:    fmt.Sprintf("%s (%s)", i.Name, i.Text),
			Percent: i.Percent,
		})
	}

	return parsedStats{
		Title: category,
		Stats: keyStatsParsed,
	}
}

func render(period string, total string, stats []parsedStats) {
	// get info for slug padding
	var maxStringLength int
	for _, stat := range stats {
		for _, i := range stat.Stats {
			slug := i.Slug
			if maxStringLength < len(slug) {
				maxStringLength = len(slug)
			}
		}
	}

	//render output
	//// print total
	// -- prettify period label -- //
	switch period {
	case "last_7_days":
		period = "Last Week"
	case "last_30_days":
		period = "Last Month"
	}
	// --------------------------- //

	color.Green(fmt.Sprintf("⏳  Total for %s", period))
	color.HiCyan(total)
	fmt.Println("")

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

			hiBlue := color.New(color.FgHiBlue).SprintFunc()
			fmt.Printf("%s : %s %v%%\n", hiBlue(slug), bar, i.Percent)
		}

		fmt.Println("")
	}
}
