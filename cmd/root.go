/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Options struct {
	Greet bool
	Lo    bool
}

var opts Options

var rootCmd = &cobra.Command{
	Use:   "softaudit",
	Short: "softaudit demo CLI",
	RunE: func(cmd *cobra.Command, args []string) error {
		if !opts.Greet && !opts.Lo {
			return cmd.Help()
		}
		if opts.Greet {
			fmt.Println("Hello from -g/--greet")
		}
		if opts.Lo {
			fmt.Println("Message from --lo")
		}
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// PersistentFlags -> доступны всем подкомандам. [web:179]
	rootCmd.PersistentFlags().BoolVarP(&opts.Greet, "greet", "g", false, "Print greeting message")
	rootCmd.PersistentFlags().BoolVar(&opts.Lo, "lo", false, "Print another message")
}
