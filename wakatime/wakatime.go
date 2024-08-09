/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package wakatime

var apiEndpoint = "https://wakatime.com"

type categoryStats []struct {
	Name    string  `json:"name"`
	Text    string  `json:"text"`
	Percent float64 `json:"percent"`
}
