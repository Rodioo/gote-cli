package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "gote",
	Version: "v1.0.0",
	Short:   "Gote is a note taking CLI tool",
}

func init() {
	rootCmd.SetVersionTemplate("Gote - {{.Version}}\n")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
