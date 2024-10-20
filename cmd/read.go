package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/antoniofalcescu/gote-cli/utils"
	"github.com/spf13/cobra"
)

// TODO: see how to implement cat or a way to preview notes (at first simple txt files)
func handleReadCmd(cmd *cobra.Command, args []string) {
	title := args[0]
	group, _ := cmd.Flags().GetString("group")

	notePath, getPathError := getNotePath(title, group)
	if getPathError != nil {
		log.Fatalf("Failed to find note: %v", getPathError)
	}

	utils.Clear()
	catCmd := exec.Command("cat", notePath)
	catCmd.Stdout = os.Stdout
	if err := catCmd.Run(); err != nil {
		log.Fatalf("Failed to read note: %v", err)
	}
}

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a note",
	Long:  "Prints the preview of your note",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:   handleReadCmd,
}

func init() {
	readCmd.Flags().StringP("group", "g", "", "The group of your note")

	rootCmd.AddCommand(readCmd)
}
