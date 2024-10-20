package cmd

import (
	"fmt"
	"io/fs"
	"path"
	"path/filepath"
	"strings"

	"github.com/antoniofalcescu/gote-cli/utils"
)

func getNotePath(title string, group string) (string, error) {
	dirPath, getStorageDirError := utils.GetStorageDirPath(osImpl)
	if getStorageDirError != nil {
		return "", getStorageDirError
	}

	matchedFiles := make([]string, 0)
	groupPath := path.Join(dirPath, group)
	walkDirError := filepath.WalkDir(groupPath, func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !dirEntry.IsDir() && strings.TrimSuffix(dirEntry.Name(), filepath.Ext(dirEntry.Name())) == title {
			matchedFiles = append(matchedFiles, path)
		}
		return nil
	})
	if walkDirError != nil {
		return "", walkDirError
	}

	if len(matchedFiles) == 0 {
		return "", fmt.Errorf("no such note with title: %s", title)
	} else if len(matchedFiles) > 1 {
		return "", fmt.Errorf("multiple notes found for title: %s\n[%s]\npass -g to specify group", title, strings.Join(matchedFiles, ", "))
	} else {
		return matchedFiles[0], nil
	}
}
