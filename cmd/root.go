package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "gote",
	Short: "Gote is a note taking CLI tool",
	Long: `Gote is a note taking CLI tool 
			which allows you to quickly create, read, edit or delete your notes directly from the terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
