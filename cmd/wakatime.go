package cmd

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

var WakatimeEndpoint = "https://wakatime.com"

// api response
type WakatimeKeyStats []struct {
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
type WakatimeStats struct {
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
		Projects         WakatimeKeyStats `json:"projects"`
		Languages        WakatimeKeyStats `json:"languages"`
		Dependencies     WakatimeKeyStats `json:"dependencies"`
		Machines         WakatimeKeyStats `json:"machines"`
		Editors          WakatimeKeyStats `json:"editors"`
		OperatingSystems WakatimeKeyStats `json:"operating_systems"`
		Categories       WakatimeKeyStats `json:"categories"`
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

type CategoryStats struct {
	Slug    string
	Percent float64
}

type Stats struct {
	Title string
	Stats []CategoryStats
}

func createAuthorizationHeader() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(viper.GetString("WAKATIME_API_KEY"))))
}

func getStats(period string) WakatimeStats {
	var response WakatimeStats
	err := requests.
		URL(WakatimeEndpoint).
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

func appendToKey(category string, keyStats WakatimeKeyStats) Stats {
	var categoryStats []CategoryStats

	for _, i := range keyStats {
		categoryStats = append(categoryStats, CategoryStats{
			Slug:    fmt.Sprintf("%s (%s)", i.Name, i.Text),
			Percent: i.Percent,
		})
	}

	return Stats{
		Title: category,
		Stats: categoryStats,
	}
}
func extractData(r WakatimeStats) (string, []Stats) {
	var total string
	var stats []Stats

	for _, i := range r.Data {
		total = i.GrandTotal.Text

		stats = append(stats, appendToKey("🚀  Projects", i.Projects))
		stats = append(stats, appendToKey("🗣️  Languages", i.Languages))
		stats = append(stats, appendToKey("✍️  Editors", i.Editors))
		stats = append(stats, appendToKey("💻  OS", i.OperatingSystems))
	}

	return total, stats
}
