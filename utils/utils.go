package utils

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
)

const APP_NAME = "Gote"

func GetStorageDirPath() (dirPath string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return dirPath, err
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		dirPath = path.Join(homeDir, "Library", "Application Support", APP_NAME)
	case "linux":
		dirPath = path.Join(homeDir, APP_NAME)
	case "windows":
		dirPath = path.Join(homeDir, fmt.Sprintf(".%s", APP_NAME))
	default:
		err = errors.New("unsupported platform, please open a github issue")
	}

	return dirPath, err
}

func CreateDirIfNotExists(dirPath string) (err error) {
	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		return nil
	}

	if err := os.Mkdir(dirPath, os.ModePerm); err != nil {
		return err
	}

	fmt.Printf("finished first time initialization, your notes will be saved in %s\n", dirPath)
	return nil
}
