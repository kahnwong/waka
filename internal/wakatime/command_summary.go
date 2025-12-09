package wakatime

import "github.com/rs/zerolog/log"

func RenderSummary(period string) {
	response, err := wakatimeClient.getSummary(period)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to get summary for %s", period)
	}
	total, stats := extractSummaryData(response)
	render(period, total, stats)
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
