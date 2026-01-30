package wakatime

import (
	"testing"
)

func TestAppendToKey(t *testing.T) {
	tests := []struct {
		name     string
		category string
		stats    categoryStats
		want     parsedStats
	}{
		{
			name:     "empty stats",
			category: "Languages",
			stats:    categoryStats{},
			want: parsedStats{
				Title: "Languages",
				Stats: []parsedCategoryStats{},
			},
		},
		{
			name:     "single stat",
			category: "Languages",
			stats: categoryStats{
				{
					Name:    "Go",
					Text:    "5 hrs 30 mins",
					Percent: 75.5,
				},
			},
			want: parsedStats{
				Title: "Languages",
				Stats: []parsedCategoryStats{
					{
						Slug:    "Go (5 hrs 30 mins)",
						Percent: 75.5,
					},
				},
			},
		},
		{
			name:     "multiple stats",
			category: "Projects",
			stats: categoryStats{
				{
					Name:    "project-a",
					Text:    "3 hrs",
					Percent: 50.0,
				},
				{
					Name:    "project-b",
					Text:    "2 hrs",
					Percent: 33.3,
				},
			},
			want: parsedStats{
				Title: "Projects",
				Stats: []parsedCategoryStats{
					{
						Slug:    "project-a (3 hrs)",
						Percent: 50.0,
					},
					{
						Slug:    "project-b (2 hrs)",
						Percent: 33.3,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := appendToKey(tt.category, tt.stats)

			if got.Title != tt.want.Title {
				t.Errorf("appendToKey() Title = %v, want %v", got.Title, tt.want.Title)
			}

			if len(got.Stats) != len(tt.want.Stats) {
				t.Errorf("appendToKey() Stats length = %v, want %v", len(got.Stats), len(tt.want.Stats))
				return
			}

			for i := range got.Stats {
				if got.Stats[i].Slug != tt.want.Stats[i].Slug {
					t.Errorf("appendToKey() Stats[%d].Slug = %v, want %v", i, got.Stats[i].Slug, tt.want.Stats[i].Slug)
				}
				if got.Stats[i].Percent != tt.want.Stats[i].Percent {
					t.Errorf("appendToKey() Stats[%d].Percent = %v, want %v", i, got.Stats[i].Percent, tt.want.Stats[i].Percent)
				}
			}
		})
	}
}
