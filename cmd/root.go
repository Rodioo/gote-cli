package cmd

import (
	"fmt"
	"os"

	"github.com/antoniofalcescu/gote-cli/utils"
	"github.com/spf13/cobra"
)

var osImpl Os

var rootCmd = &cobra.Command{
	Use:     "gote",
	Version: utils.Version,
	Short:   "Gote is a note taking CLI tool",
}

func init() {
	osImpl = Os{}
	rootCmd.SetVersionTemplate("Gote - {{.Version}}\n")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
