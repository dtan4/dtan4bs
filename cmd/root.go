// Copyright © 2017 NAME HERE <EMAIL ADDRESS>

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dtan4bs",
	Short: "Project skeleton generator for @dtan4",
}

var rootOpts struct {
	debug bool
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		if rootOpts.debug {
			fmt.Fprintf(os.Stderr, "%+v\n", err)
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().BoolVar(&rootOpts.debug, "debug", false, "debug mode")
}
