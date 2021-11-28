package util

import (
	"errors"
	"os"
)

func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}
