package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func handleDeleteCommand(cmd *cobra.Command, args []string) {
	title := args[0]
	group, _ := cmd.Flags().GetString("group")

	notePath, getPathError := getNotePath(title, group)
	if getPathError != nil {
		log.Fatalf("Failed to find note: %v", getPathError)
	}

	fmt.Printf("Are you sure you want to delete note %s?[y/n]", title)
	reader := bufio.NewReader(os.Stdin)
	deleteConfirmation, _, deleteConfirmationErr := reader.ReadRune()

	if deleteConfirmationErr != nil {
		log.Fatalf("Unexpected error %v", deleteConfirmationErr)
	}

	if deleteConfirmation == 'y' || deleteConfirmation == 'Y' {
		rmCmd := exec.Command("rm", notePath)
		rmCmd.Stdout = os.Stdout
		if err := rmCmd.Run(); err != nil {
			log.Fatalf("Failed to delete note: %v", err)
		}
		fmt.Printf("Note %s deleted succesfully\n", title)
	} else if deleteConfirmation == 'n' || deleteConfirmation == 'N' {
		return
	} else {
		log.Fatalf("Invalid response %s", string(deleteConfirmation))
	}
}

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a note",
	Long:  "Deletes a note based on title",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:   handleDeleteCommand,
}

func init() {
	deleteCmd.Flags().StringP("group", "g", "", "The group you want to delete")

	rootCmd.AddCommand(deleteCmd)
}
