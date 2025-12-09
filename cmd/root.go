/*
Copyright © 2024 Karn Wong <karn@karnwong.me>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
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

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
