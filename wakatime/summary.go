/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package wakatime

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/fatih/color"

	"github.com/carlmjohnson/requests"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// api response
type summaryResponse struct {
	Data []struct {
		GrandTotal struct {
			Text string `json:"text"`
		} `json:"grand_total"`
		Projects         categoryStats `json:"projects"`
		Languages        categoryStats `json:"languages"`
		Dependencies     categoryStats `json:"dependencies"`
		Machines         categoryStats `json:"machines"`
		Editors          categoryStats `json:"editors"`
		OperatingSystems categoryStats `json:"operating_systems"`
		Categories       categoryStats `json:"categories"`
	} `json:"data"`
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

// for parsing
type parsedCategoryStats struct {
	Slug    string
	Percent float64
}

type parsedStats struct {
	Title string
	Stats []parsedCategoryStats
}

// main
func createAuthorizationHeader() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(viper.GetString("WAKATIME_API_KEY"))))
}

func getSummary(period string) summaryResponse {
	var response summaryResponse
	err := requests.
		URL(apiEndpoint).
		Method(http.MethodGet).
		Path("api/v1/users/current/summaries").
		Param("range", period).
		Header("Authorization", createAuthorizationHeader()).
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to get summary for %s", period)
	}

	return response
}

func appendToKey(category string, keyStats categoryStats) parsedStats {
	var categoryStats []parsedCategoryStats

	for _, i := range keyStats {
		categoryStats = append(categoryStats, parsedCategoryStats{
			Slug:    fmt.Sprintf("%s (%s)", i.Name, i.Text),
			Percent: i.Percent,
		})
	}

	return parsedStats{
		Title: category,
		Stats: categoryStats,
	}
}
func extractData(r summaryResponse) (string, []parsedStats) {
	var total string
	var stats []parsedStats

	for _, i := range r.Data {
		total = i.GrandTotal.Text

		stats = append(stats, appendToKey("💻  OS", i.OperatingSystems))
		stats = append(stats, appendToKey("✍️  Editors", i.Editors))
		stats = append(stats, appendToKey("🗣️  Languages", i.Languages))
		stats = append(stats, appendToKey("🚀  Projects", i.Projects))
	}

	return total, stats
}

func Render(period string) {
	response := getSummary(period)
	total, stats := extractData(response)

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
	//// print total
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
