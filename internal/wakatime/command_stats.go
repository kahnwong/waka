package wakatime

import (
	"github.com/rs/zerolog/log"
)

type parsedStats struct {
	Title string
	Stats []parsedCategoryStats
}

func RenderStats(period string) {
	response, err := wakatimeClient.getStats(period)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to get stats for %s", period)
	}
	total, stats := extractStatsData(response)
	render(period, total, stats)
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
