package util

import (
	"errors"
	"os"
	"strings"
)

func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); errors.Is(err, os.ErrNotExist) {
		return false
	}

	return true
}

func AllowedFormats(format string) bool {
	switch strings.ToLower(format) {
	case ".jpg", "jpeg", "png":
		return true
	default:
		return false
	}
}
