package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/antoniofalcescu/gote-cli/types"
)

func GetStorageDirPath(osProvider types.OsProvider) (dirPath string, err error) {
	homeDir, err := osProvider.GetHomeDir()
	if err != nil {
		return dirPath, err
	}

	switch os := osProvider.GetOs(); os {
	case "darwin":
		dirPath = path.Join(homeDir, "Library", "Application Support", AppName)
	case "linux":
		dirPath = path.Join(homeDir, AppName)
	case "windows":
		dirPath = path.Join(homeDir, fmt.Sprintf(".%s", AppName))
	default:
		err = fmt.Errorf("unsupported platform: %s, please open a github issue", os)
	}

	return dirPath, err
}

func CreateDirIfNotExists(dirPath string) error {

	if _, err := os.Stat(dirPath); !os.IsNotExist(err) {
		return nil
	}

	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	return nil
}

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func Clear() error {
	var clearCmd *exec.Cmd

	switch os := runtime.GOOS; os {
	case "darwin":
		clearCmd = exec.Command("clear")
	case "linux":
		clearCmd = exec.Command("clear")
	case "windows":
		clearCmd = exec.Command("cmd", "/c", "cls")
	default:
		return fmt.Errorf("unsupported platform: %s, please open a github issue", os)
	}

	clearCmd.Stdout = os.Stdout
	clearCmd.Run()
	return nil
}
