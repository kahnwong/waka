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

type StatsResponse struct {
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
