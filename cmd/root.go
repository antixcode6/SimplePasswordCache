package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "spwc",
	Short: "Simple Password Cache is a Go re-imagining of pass(1)",
	Long:  `A Go re-imagining of pass(1) built with love by me for anyone who wants to use it.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
