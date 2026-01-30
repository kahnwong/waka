package wakatime

import (
	"testing"
	"time"
)

func TestExtractSummaryData(t *testing.T) {
	tests := []struct {
		name         string
		response     summaryResponse
		wantTotal    string
		wantNumStats int
	}{
		{
			name: "valid summary response",
			response: summaryResponse{
				Start: time.Now(),
				End:   time.Now(),
				Data: []struct {
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
				}{
					{
						GrandTotal: struct {
							Text string `json:"text"`
						}{
							Text: "8 hrs 15 mins",
						},
						OperatingSystems: categoryStats{
							{Name: "macOS", Text: "8 hrs", Percent: 100.0},
						},
						Editors: categoryStats{
							{Name: "Vim", Text: "8 hrs", Percent: 100.0},
						},
						Languages: categoryStats{
							{Name: "Go", Text: "6 hrs", Percent: 75.0},
							{Name: "YAML", Text: "2 hrs", Percent: 25.0},
						},
						Projects: categoryStats{
							{Name: "utils", Text: "8 hrs", Percent: 100.0},
						},
					},
				},
			},
			wantTotal:    "8 hrs 15 mins",
			wantNumStats: 4, // OS, Editors, Languages, Projects
		},
		{
			name: "empty summary response",
			response: summaryResponse{
				Start: time.Now(),
				End:   time.Now(),
				Data: []struct {
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
				}{
					{
						GrandTotal: struct {
							Text string `json:"text"`
						}{
							Text: "0 secs",
						},
						OperatingSystems: categoryStats{},
						Editors:          categoryStats{},
						Languages:        categoryStats{},
						Projects:         categoryStats{},
					},
				},
			},
			wantTotal:    "0 secs",
			wantNumStats: 4,
		},
		{
			name: "no data in response",
			response: summaryResponse{
				Start: time.Now(),
				End:   time.Now(),
				Data: []struct {
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
				}{},
			},
			wantTotal:    "",
			wantNumStats: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTotal, gotStats := extractSummaryData(tt.response)

			if gotTotal != tt.wantTotal {
				t.Errorf("extractSummaryData() total = %v, want %v", gotTotal, tt.wantTotal)
			}

			if len(gotStats) != tt.wantNumStats {
				t.Errorf("extractSummaryData() stats length = %v, want %v", len(gotStats), tt.wantNumStats)
			}

			// Verify the order and titles for non-empty responses
			if tt.wantNumStats > 0 {
				expectedTitles := []string{"💻  OS", "✍️  Editors", "🗣️  Languages", "🚀  Projects"}
				for i, title := range expectedTitles {
					if i < len(gotStats) && gotStats[i].Title != title {
						t.Errorf("extractSummaryData() stats[%d].Title = %v, want %v", i, gotStats[i].Title, title)
					}
				}
			}
		})
	}
}
