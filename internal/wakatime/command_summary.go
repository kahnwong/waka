package wakatime

import "fmt"

func RenderSummary(period string) error {
	if err := ensureInitialized(); err != nil {
		return fmt.Errorf("failed to initialize: %w", err)
	}

	response, err := wakatimeClient.getSummary(period)
	if err != nil {
		return fmt.Errorf("failed to get summary for %s: %w", period, err)
	}
	total, stats := extractSummaryData(response)
	render(period, total, stats)
	return nil
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
