/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	//Run: func(cmd *cobra.Command, args []string) {
	//},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// read config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.config/waka")

	err := viper.ReadInConfig()
	if err != nil {
		// get API key from user input
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Wakatime API Key: ")

		apiKey, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal().Err(err).Msg("Error reading user input")
		}

		// write config
		//// create config dir
		homeDir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal().Err(err).Msg("Error obtaining home directory")
		}
		err = os.MkdirAll(filepath.Join(homeDir, ".config", "waka"), os.ModePerm)
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating config path")
		}

		//// write yaml
		filename := filepath.Join(homeDir, ".config", "waka", "config.yaml")
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
		if err != nil {
			log.Fatal().Err(err).Msg("Error creating config file")
		}
		defer file.Close()

		enc := yaml.NewEncoder(file)
		err = enc.Encode(Config{WakatimeApiKey: strings.TrimSpace(apiKey)})

		if err != nil {
			log.Fatal().Err(err).Msg("Error writing config")
		}

		// read config to viper
		err = viper.ReadInConfig()
		if err != nil {
			log.Fatal().Err(err).Msg("Error reading config after config initialization")
		}

	}

	// rootCmd
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
