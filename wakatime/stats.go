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

type statsCategoryStats []struct {
	TotalSeconds  float64 `json:"total_seconds"`
	Name          string  `json:"name"`
	MachineNameID string  `json:"machine_name_id"`
	Percent       float64 `json:"percent"`
	Digital       string  `json:"digital"`
	Decimal       string  `json:"decimal"`
	Text          string  `json:"text"`
	Hours         int     `json:"hours"`
	Minutes       int     `json:"minutes"`
}
type StatsResponse struct {
	Data struct {
		ID                                              string             `json:"id"`
		UserID                                          string             `json:"user_id"`
		Range                                           string             `json:"range"`
		Start                                           time.Time          `json:"start"`
		End                                             time.Time          `json:"end"`
		Timeout                                         int                `json:"timeout"`
		WritesOnly                                      bool               `json:"writes_only"`
		Timezone                                        string             `json:"timezone"`
		Holidays                                        int                `json:"holidays"`
		Status                                          string             `json:"status"`
		CreatedAt                                       time.Time          `json:"created_at"`
		ModifiedAt                                      time.Time          `json:"modified_at"`
		Machines                                        statsCategoryStats `json:"machines"`
		DailyAverageIncludingOtherLanguage              float64            `json:"daily_average_including_other_language"`
		TotalSeconds                                    float64            `json:"total_seconds"`
		Projects                                        statsCategoryStats `json:"projects"`
		PercentCalculated                               int                `json:"percent_calculated"`
		HumanReadableTotalIncludingOtherLanguage        string             `json:"human_readable_total_including_other_language"`
		HumanReadableDailyAverage                       string             `json:"human_readable_daily_average"`
		HumanReadableDailyAverageIncludingOtherLanguage string             `json:"human_readable_daily_average_including_other_language"`
		Editors                                         statsCategoryStats `json:"editors"`
		OperatingSystems                                statsCategoryStats `json:"operating_systems"`
		TotalSecondsIncludingOtherLanguage              float64            `json:"total_seconds_including_other_language"`
		DailyAverage                                    float64            `json:"daily_average"`
		IsAlreadyUpdating                               bool               `json:"is_already_updating"`
		IsUpToDate                                      bool               `json:"is_up_to_date"`
		HumanReadableTotal                              string             `json:"human_readable_total"`
		DaysMinusHolidays                               int                `json:"days_minus_holidays"`
		Dependencies                                    statsCategoryStats `json:"dependencies"`
		IsUpToDatePendingFuture                         bool               `json:"is_up_to_date_pending_future"`
		IsStuck                                         bool               `json:"is_stuck"`
		Categories                                      statsCategoryStats `json:"categories"`
		BestDay                                         struct {
			Date         string  `json:"date"`
			TotalSeconds float64 `json:"total_seconds"`
			Text         string  `json:"text"`
		} `json:"best_day"`
		Languages               statsCategoryStats `json:"languages"`
		DaysIncludingHolidays   int                `json:"days_including_holidays"`
		IsCached                bool               `json:"is_cached"`
		Username                interface{}        `json:"username"`
		IsIncludingToday        bool               `json:"is_including_today"`
		HumanReadableRange      string             `json:"human_readable_range"`
		IsCodingActivityVisible bool               `json:"is_coding_activity_visible"`
		IsOtherUsageVisible     bool               `json:"is_other_usage_visible"`
	} `json:"data"`
}

func GetStats(period string) StatsResponse {
	var response StatsResponse
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
