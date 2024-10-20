package cmd

import (
	"github.com/spf13/cobra"
)

func handleVersionCmd(cmd *cobra.Command, args []string) {
	root := cmd.Root()
	root.SetArgs([]string{"--version"})
	root.Execute()
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Current version",
	Long:  "The current version of the gote CLI",
	Run:   handleVersionCmd,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
