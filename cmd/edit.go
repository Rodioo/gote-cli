package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func handleEditCommand(cmd *cobra.Command, args []string) {
	title := args[0]
	group, _ := cmd.Flags().GetString("group")

	notePath, getPathError := getNotePath(title, group)
	if getPathError != nil {
		log.Fatalf("Failed to find note: %v", getPathError)
	}

	vimCmd := exec.Command("vim", notePath)
	vimCmd.Stdout = os.Stdout
	vimCmd.Stdin = os.Stdin
	vimCmd.Stderr = os.Stderr

	if err := vimCmd.Run(); err != nil {
		log.Fatalf("Failed to open Vim: %v", err)
	}

	fmt.Printf("Note %s edited succesfully\n", title)
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a note",
	Long:  "Opens the note to edit inside the terminal",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:   handleEditCommand,
}

func init() {
	editCmd.Flags().StringP("group", "g", "", "The group of your note")

	rootCmd.AddCommand(editCmd)
}
