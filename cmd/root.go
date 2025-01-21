/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/kahnwong/waka/wakatime"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	version = "dev"
)

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
	err = enc.Encode(wakatime.Config{WakatimeApiKey: strings.TrimSpace(apiKey)})

	if err != nil {
		log.Fatal().Msg("Error writing config")
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
