package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"slices"
	"strings"

	"github.com/antoniofalcescu/gote-cli/utils"
	"github.com/spf13/cobra"
)

func validateFlags(group string, title string, format string) error {
	// TODO: length should be configurable
	if length := len(group); length > 32 {
		return fmt.Errorf("group name too long: got %d, max is 32", length)
	}

	if length := len(title); length > 32 {
		return fmt.Errorf("title name too long: got %d, max is 32", length)
	}

	if allowedFormats := utils.GetAllowedFormats(); !slices.Contains(allowedFormats, format) {
		return fmt.Errorf("invalid format: got %s, allowed formats are [%s]", format, strings.Join(allowedFormats, ", "))
	}

	return nil
}

func validateUniqueNote(groupPath string, title string, group string) error {
	return filepath.WalkDir(groupPath, func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !dirEntry.IsDir() && strings.TrimSuffix(dirEntry.Name(), filepath.Ext(dirEntry.Name())) == title {
			return fmt.Errorf("note with title %s already exists in group %s", title, group)
		}
		return nil
	})
}

func handleAddCmd(cmd *cobra.Command, args []string) {
	title := args[0]
	group, _ := cmd.Flags().GetString("group")
	format, _ := cmd.Flags().GetString("format")

	if err := validateFlags(group, title, format); err != nil {
		log.Fatal(err)
	}

	dirPath, err := utils.GetStorageDirPath(osImpl)
	if err != nil {
		log.Fatal(err)
	}

	groupPath := path.Join(dirPath, group)

	if err := utils.CreateDirIfNotExists(groupPath); err != nil {
		log.Fatal(err)
	}

	if err := validateUniqueNote(groupPath, title, group); err != nil {
		log.Fatal(err)
	}

	notePath := path.Join(groupPath, fmt.Sprintf("%s.%s", title, format))

	vimCmd := exec.Command("vim", notePath)
	vimCmd.Stdout = os.Stdout
	vimCmd.Stdin = os.Stdin
	vimCmd.Stderr = os.Stderr

	if err := vimCmd.Run(); err != nil {
		log.Fatalf("Failed to open Vim: %v", err)
	}

	fmt.Printf("Note %s added succesfully\n", title)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new note",
	Long:  "Opens a new file inside the terminal to add a note",
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:   handleAddCmd,
}

func init() {
	addCmd.Flags().StringP("group", "g", "general", "The group of notes in which you want to add")
	addCmd.Flags().StringP("format", "f", "txt", "The format of the note's file")

	rootCmd.AddCommand(addCmd)
}
