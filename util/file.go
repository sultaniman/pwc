package util

import (
	"errors"
	"os"
	"path/filepath"
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

func WalkAndReadDirectory(path string) ([]string, error) {
	var images []string
	err := filepath.Walk(path, func(currentPath string, info os.FileInfo, err error) error {
		if mode := info.Mode(); mode.IsRegular() {
			if AllowedFormats(filepath.Ext(currentPath)) {
				images = append(images, currentPath)
			}
		}

		return err
	})

	if err != nil {
		return nil, err
	}

	return images, nil
}
