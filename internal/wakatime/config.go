package wakatime

import (
	"strings"

	"gopkg.in/ini.v1"
)

const defaultAPIURL = "https://wakatime.com/api"

var configPath = "~/.wakatime.cfg"

type Config struct {
	APIURL string
	APIKey string
}

func readConfig(path string) (Config, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return Config{}, err
	}

	settings := cfg.Section("settings")
	apiURL := settings.Key("api_url").String()
	if apiURL == "" {
		apiURL = defaultAPIURL
	} else {
		apiURL = normalizeAPIURL(apiURL)
	}

	return Config{
		APIURL: apiURL,
		APIKey: settings.Key("api_key").String(),
	}, nil
}

func normalizeAPIURL(apiURL string) string {
	apiURL = strings.TrimRight(apiURL, "/")
	if strings.Contains(apiURL, "wakapi") && !strings.HasSuffix(apiURL, "/compat/wakatime/v1") {
		return apiURL + "/compat/wakatime/v1/"
	}
	if !strings.HasSuffix(apiURL, "/v1") {
		return apiURL + "/v1/"
	}
	return apiURL + "/"
}
