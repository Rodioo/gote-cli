package utils

import (
	"errors"
	"runtime"
)

func GetPathByOs() (path string, err error) {
	switch os := runtime.GOOS; os {
	case "darwin":
		path = "darwin"
	case "linux":
		path = "linux"
	case "windows":
		path = "windows"
	default:
		err = errors.New("unsupported platform, please open a github issue")
	}

	return path, err
}
