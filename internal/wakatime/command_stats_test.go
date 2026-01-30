package wakatime

import (
	"testing"
	"time"
)

func TestExtractStatsData(t *testing.T) {
	tests := []struct {
		name         string
		response     statsResponse
		wantTotal    string
		wantNumStats int
	}{
		{
			name: "valid stats response",
			response: statsResponse{
				Data: struct {
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
				}{
					HumanReadableTotal: "10 hrs 30 mins",
					OperatingSystems: categoryStats{
						{Name: "Linux", Text: "8 hrs", Percent: 80.0},
					},
					Editors: categoryStats{
						{Name: "VS Code", Text: "10 hrs", Percent: 100.0},
					},
					Languages: categoryStats{
						{Name: "Go", Text: "5 hrs", Percent: 50.0},
						{Name: "Python", Text: "3 hrs", Percent: 30.0},
					},
					Projects: categoryStats{
						{Name: "waka", Text: "10 hrs", Percent: 100.0},
					},
				},
			},
			wantTotal:    "10 hrs 30 mins",
			wantNumStats: 4, // OS, Editors, Languages, Projects
		},
		{
			name: "empty stats response",
			response: statsResponse{
				Data: struct {
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
				}{
					HumanReadableTotal: "0 secs",
					OperatingSystems:   categoryStats{},
					Editors:            categoryStats{},
					Languages:          categoryStats{},
					Projects:           categoryStats{},
				},
			},
			wantTotal:    "0 secs",
			wantNumStats: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTotal, gotStats := extractStatsData(tt.response)

			if gotTotal != tt.wantTotal {
				t.Errorf("extractStatsData() total = %v, want %v", gotTotal, tt.wantTotal)
			}

			if len(gotStats) != tt.wantNumStats {
				t.Errorf("extractStatsData() stats length = %v, want %v", len(gotStats), tt.wantNumStats)
			}

			// Verify the order and titles
			expectedTitles := []string{"💻  OS", "✍️  Editors", "🗣️  Languages", "🚀  Projects"}
			for i, title := range expectedTitles {
				if i < len(gotStats) && gotStats[i].Title != title {
					t.Errorf("extractStatsData() stats[%d].Title = %v, want %v", i, gotStats[i].Title, title)
				}
			}
		})
	}
}
