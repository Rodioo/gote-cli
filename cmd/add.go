package cmd

import (
	"log"

	"github.com/antoniofalcescu/gote-cli/utils"
	"github.com/spf13/cobra"
)

func handleAddCmd(cmd *cobra.Command, args []string) {
	dirPath, err := utils.GetStorageDirPath()

	if err != nil {
		log.Fatal(err)
	}

	if err := utils.CreateDirIfNotExists(dirPath); err != nil {
		log.Fatal(err)
	}
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new note",
	Long:  "Opens a new MD file inside the terminal to add a note",
	Run:   handleAddCmd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}
