/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	cliBase "github.com/kahnwong/cli-base"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

var (
	version = "dev"
)

type Config struct {
	WakatimeApiKey string `yaml:"WAKATIME_API_KEY"`
}

var rootCmd = &cobra.Command{
	Use:     "waka",
	Version: version,
	Short:   "Display wakatime stats in your terminal",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func InitConfigIfNotExists(path string) {
	// get API key from user input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Wakatime API Key: ")

	apiKey, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal().Msg("Error reading user input")
	}

	// write to yaml
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal().Msg("Error creating config file")
	}
	defer file.Close()

	enc := yaml.NewEncoder(file)
	err = enc.Encode(Config{WakatimeApiKey: strings.TrimSpace(apiKey)})

	if err != nil {
		log.Fatal().Msg("Error writing config")
	}
}

func init() {
	// init config if does not exists
	path := "~/.config/waka/config.yaml"
	path, err := cliBase.CheckIfConfigExists(path)
	if err != nil {
		InitConfigIfNotExists(path)
		log.Info().Msg("Successfully initialized config")
	}

	// read config
	config := cliBase.ReadYaml[Config](path)
	viper.SetDefault("WAKATIME_API_KEY", config.WakatimeApiKey)

	// rootCmd
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
