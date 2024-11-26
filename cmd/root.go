/*
Copyright © 2024 alexeev-prog
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "carbon_pkg",
	Short: "Blazing fast and modern package manager written in Golang",
	Long: `carbonpkg is a package manager designed to simplify the process of installing software hosted on GitHub.

It provides users with a convenient interface to install, update, and remove packages, as well as manage dependencies.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.carbon_pkg.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
