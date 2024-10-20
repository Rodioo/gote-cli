package cmd

import (
	"os"
	"runtime"
)

type Os struct{}

func (osImpl Os) GetOs() string {
	return runtime.GOOS
}

func (osImpl Os) GetHomeDir() (string, error) {
	return os.UserHomeDir()
}
