package cmd

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var WakatimeEndpoint = "https://wakatime.com"

type WakatimeStats struct {
	Data []struct {
		GrandTotal struct {
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			TotalSeconds float64 `json:"total_seconds"`
			Digital      string  `json:"digital"`
			Decimal      string  `json:"decimal"`
			Text         string  `json:"text"`
		} `json:"grand_total"`
		Range struct {
			Start    time.Time `json:"start"`
			End      time.Time `json:"end"`
			Date     string    `json:"date"`
			Text     string    `json:"text"`
			Timezone string    `json:"timezone"`
		} `json:"range"`
		Projects []struct {
			Name         string      `json:"name"`
			TotalSeconds float64     `json:"total_seconds"`
			Color        interface{} `json:"color"`
			Digital      string      `json:"digital"`
			Decimal      string      `json:"decimal"`
			Text         string      `json:"text"`
			Hours        int         `json:"hours"`
			Minutes      int         `json:"minutes"`
			Seconds      int         `json:"seconds"`
			Percent      float64     `json:"percent"`
		} `json:"projects"`
		Languages []struct {
			Name         string  `json:"name"`
			TotalSeconds float64 `json:"total_seconds"`
			Digital      string  `json:"digital"`
			Decimal      string  `json:"decimal"`
			Text         string  `json:"text"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Seconds      int     `json:"seconds"`
			Percent      float64 `json:"percent"`
		} `json:"languages"`
		Dependencies []struct {
			Name         string  `json:"name"`
			TotalSeconds float64 `json:"total_seconds"`
			Digital      string  `json:"digital"`
			Decimal      string  `json:"decimal"`
			Text         string  `json:"text"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Seconds      int     `json:"seconds"`
			Percent      float64 `json:"percent"`
		} `json:"dependencies"`
		Machines []struct {
			Name          string  `json:"name"`
			TotalSeconds  float64 `json:"total_seconds"`
			MachineNameID string  `json:"machine_name_id"`
			Digital       string  `json:"digital"`
			Decimal       string  `json:"decimal"`
			Text          string  `json:"text"`
			Hours         int     `json:"hours"`
			Minutes       int     `json:"minutes"`
			Seconds       int     `json:"seconds"`
			Percent       float64 `json:"percent"`
		} `json:"machines"`
		Editors []struct {
			Name         string  `json:"name"`
			TotalSeconds float64 `json:"total_seconds"`
			Digital      string  `json:"digital"`
			Decimal      string  `json:"decimal"`
			Text         string  `json:"text"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Seconds      int     `json:"seconds"`
			Percent      float64 `json:"percent"`
		} `json:"editors"`
		OperatingSystems []struct {
			Name         string  `json:"name"`
			TotalSeconds float64 `json:"total_seconds"`
			Digital      string  `json:"digital"`
			Decimal      string  `json:"decimal"`
			Text         string  `json:"text"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Seconds      int     `json:"seconds"`
			Percent      float64 `json:"percent"`
		} `json:"operating_systems"`
		Categories []struct {
			Name         string  `json:"name"`
			TotalSeconds float64 `json:"total_seconds"`
			Digital      string  `json:"digital"`
			Decimal      string  `json:"decimal"`
			Text         string  `json:"text"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Seconds      int     `json:"seconds"`
			Percent      float64 `json:"percent"`
		} `json:"categories"`
	} `json:"data"`
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	CumulativeTotal struct {
		Seconds float64 `json:"seconds"`
		Text    string  `json:"text"`
		Digital string  `json:"digital"`
		Decimal string  `json:"decimal"`
	} `json:"cumulative_total"`
	DailyAverage struct {
		Holidays                      int    `json:"holidays"`
		DaysMinusHolidays             int    `json:"days_minus_holidays"`
		DaysIncludingHolidays         int    `json:"days_including_holidays"`
		Seconds                       int    `json:"seconds"`
		SecondsIncludingOtherLanguage int    `json:"seconds_including_other_language"`
		Text                          string `json:"text"`
		TextIncludingOtherLanguage    string `json:"text_including_other_language"`
	} `json:"daily_average"`
}

type Stats struct {
	Slug    string
	Percent float64
}

func createAuthorizationHeader() string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(viper.GetString("WAKATIME_API_KEY"))))
}
