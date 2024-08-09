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

type statsResponse struct {
	Data struct {
		Start              time.Time     `json:"start"`
		End                time.Time     `json:"end"`
		HumanReadableTotal string        `json:"human_readable_total"`
		Machines           categoryStats `json:"machines"`
		Projects           categoryStats `json:"projects"`
		Editors            categoryStats `json:"editors"`
		OperatingSystems   categoryStats `json:"operating_systems"`
		Dependencies       categoryStats `json:"dependencies"`
		Categories         categoryStats `json:"categories"`
		Languages          categoryStats `json:"languages"`
	} `json:"data"`
}

func getStats(period string) statsResponse {
	var response statsResponse
	err := requests.
		URL(apiEndpoint).
		Method(http.MethodGet).
		Pathf("api/v1/users/current/stats/%s", period).
		Header("Authorization", createAuthorizationHeader()).
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to get stats for %s", period)
	}

	return response
}

func extractStatsData(r statsResponse) (string, []parsedStats) {
	var total string
	var stats []parsedStats

	total = r.Data.HumanReadableTotal

	stats = append(stats, appendToKey("💻  OS", r.Data.OperatingSystems))
	stats = append(stats, appendToKey("✍️  Editors", r.Data.Editors))
	stats = append(stats, appendToKey("🗣️  Languages", r.Data.Languages))
	stats = append(stats, appendToKey("🚀  Projects", r.Data.Projects))

	return total, stats
}

func RenderStats(period string) {
	response := getStats(period)
	total, stats := extractStatsData(response)

	render(period, total, stats)
}
