/*
Copyright © 2026 Karn Wong <karn@karnwong.me>
*/
package wakatime

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/carlmjohnson/requests"
	cliBase "github.com/kahnwong/cli-base"
)

var wakatimeClient *Client
var initOnce sync.Once
var initErr error

type Client struct {
	baseURL             string
	client              *http.Client
	authorizationHeader string
}

type categoryStats []struct {
	Name    string  `json:"name"`
	Text    string  `json:"text"`
	Percent float64 `json:"percent"`
}

type statsResponse struct {
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

type summaryResponse struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Data  []struct {
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
	} `json:"data"`
}

func NewClient(apiKey string, apiURL string) (*Client, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("API key cannot be empty")
	}
	if apiURL == "" {
		apiURL = defaultAPIURL
	}

	c := &Client{
		baseURL:             normalizeAPIURL(apiURL),
		client:              &http.Client{},
		authorizationHeader: fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(apiKey))),
	}

	return c, nil
}

func (c *Client) getStats(period string) (statsResponse, error) {
	var response statsResponse
	err := requests.
		URL(c.baseURL).
		Method(http.MethodGet).
		Pathf("users/current/stats/%s", period).
		Header("Authorization", c.authorizationHeader).
		Client(c.client).
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		return response, fmt.Errorf("failed to get stats for %s: %w", period, err)
	}

	return response, nil
}

func (c *Client) getSummary(period string) (summaryResponse, error) {
	var response summaryResponse
	err := requests.
		URL(c.baseURL).
		Method(http.MethodGet).
		Path("users/current/summaries").
		Param("range", period).
		Header("Authorization", c.authorizationHeader).
		Client(c.client).
		ToJSON(&response).
		Fetch(context.Background())
	if err != nil {
		return response, fmt.Errorf("failed to get summary for %s: %w", period, err)
	}

	return response, nil
}

func initialize() error {
	path, err := cliBase.ExpandHome(configPath)
	if err != nil {
		return fmt.Errorf("failed to expand config path: %w", err)
	}

	config, err := readConfig(path)
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	wakatimeClient, err = NewClient(config.APIKey, config.APIURL)
	if err != nil {
		return fmt.Errorf("failed to create wakatime client: %w", err)
	}

	return nil
}

func ensureInitialized() error {
	initOnce.Do(func() {
		initErr = initialize()
	})
	return initErr
}
