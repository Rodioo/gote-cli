package utils_test

import (
	"testing"

	"github.com/antoniofalcescu/gote-cli/utils"
)

type MockMacOs struct{}

func (macOs MockMacOs) GetOs() string {
	return "darwin"
}

func (macOs MockMacOs) GetHomeDir() (string, error) {
	return "/Users/test_user", nil
}

// TODO: refactor all erors from %s to %v or %w
func TestMacStoragePath(t *testing.T) {
	macOs := MockMacOs{}
	const expectedResult = "/Users/test_user/Library/Application Support/Gote"

	result, err := utils.GetStorageDirPath(macOs)

	if result != expectedResult || err != nil {
		t.Fatalf("TestMacStoragePath() - got %s, %v, expected %s, nil", result, err, expectedResult)
	}
}
