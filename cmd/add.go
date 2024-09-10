package cmd

import (
	"fmt"
	"log"

	"github.com/antoniofalcescu/gote-cli/utils"
	"github.com/spf13/cobra"
)

func handleAddCmd(cmd *cobra.Command, args []string) {
	path, err := utils.GetPathByOs()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(path)
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
