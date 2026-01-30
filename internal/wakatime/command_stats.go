package wakatime

import (
	"fmt"
)

type parsedStats struct {
	Title string
	Stats []parsedCategoryStats
}

func RenderStats(period string) error {
	if err := ensureInitialized(); err != nil {
		return fmt.Errorf("failed to initialize: %w", err)
	}

	response, err := wakatimeClient.getStats(period)
	if err != nil {
		return fmt.Errorf("failed to get stats for %s: %w", period, err)
	}
	total, stats := extractStatsData(response)
	render(period, total, stats)
	return nil
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
