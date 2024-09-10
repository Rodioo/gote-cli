package cmd

import (
	"github.com/spf13/cobra"
)

func printVersion(cmd *cobra.Command, args []string) {
	root := cmd.Root()
	root.SetArgs([]string{"--version"})
	root.Execute()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Current version",
	Long: "The current version of the gote CLI",
	Run: printVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}