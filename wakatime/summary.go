/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package wakatime

import (
	"context"
	"net/http"
	"time"

	"github.com/carlmjohnson/requests"
	"github.com/rs/zerolog/log"
)

type summaryResponse struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Data  []struct {
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
		log.Fatal().Msgf("Failed to get summary for %s", period)
	}

	return response
}

func extractSummaryData(r summaryResponse) (string, []parsedStats) {
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

func RenderSummary(period string) {
	response := getSummary(period)
	total, stats := extractSummaryData(response)

	render(period, total, stats)
}
