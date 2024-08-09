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

var apiEndpoint = "https://wakatime.com"

// api response
type summaryCategoryStats []struct {
	Name         string      `json:"name"`
	TotalSeconds float64     `json:"total_seconds"`
	Color        interface{} `json:"color"`
	Digital      string      `json:"digital"`
	Decimal      string      `json:"decimal"`
	Text         string      `json:"text"`
	Hours        int         `json:"hours"`
	Minutes      int         `json:"minutes"`
	Seconds      int         `json:"seconds"`
	Percent      float64     `json:"percent"`
}
type summary struct {
	Data []struct {
		GrandTotal struct {
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			TotalSeconds float64 `json:"total_seconds"`
			Digital      string  `json:"digital"`
			Decimal      string  `json:"decimal"`
			Text         string  `json:"text"`
		} `json:"grand_total"`
		Range struct {
			Start    time.Time `json:"start"`
			End      time.Time `json:"end"`
			Date     string    `json:"date"`
			Text     string    `json:"text"`
			Timezone string    `json:"timezone"`
		} `json:"range"`
		Projects         summaryCategoryStats `json:"projects"`
		Languages        summaryCategoryStats `json:"languages"`
		Dependencies     summaryCategoryStats `json:"dependencies"`
		Machines         summaryCategoryStats `json:"machines"`
		Editors          summaryCategoryStats `json:"editors"`
		OperatingSystems summaryCategoryStats `json:"operating_systems"`
		Categories       summaryCategoryStats `json:"categories"`
	} `json:"data"`
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	CumulativeTotal struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
		Digital string  `json:"digital"`
		Decimal string  `json:"decimal"`
	} `json:"cumulative_total"`
	DailyAverage struct {
		Holidays                      int    `json:"holidays"`
		DaysMinusHolidays             int    `json:"days_minus_holidays"`
		DaysIncludingHolidays         int    `json:"days_including_holidays"`
		Seconds                       int    `json:"seconds"`
		SecondsIncludingOtherLanguage int    `json:"seconds_including_other_language"`
		Text                          string `json:"text"`
		TextIncludingOtherLanguage    string `json:"text_including_other_language"`
	} `json:"daily_average"`
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

func createAuthorizationHeader() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(viper.GetString("WAKATIME_API_KEY"))))
}

func getStats(period string) summary {
	var response summary
	err := requests.
		URL(apiEndpoint).
		Method(http.MethodGet).
		Path("api/v1/users/current/summaries").
		Param("range", period).
		Header("Authorization", createAuthorizationHeader()).
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get stats for today")
	}

	return response
}

func appendToKey(category string, keyStats summaryCategoryStats) parsedStats {
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
func extractData(r summary) (string, []parsedStats) {
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

func RenderStats(period string) {
	response := getStats(period)
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
