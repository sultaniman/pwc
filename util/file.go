package util

import (
	"errors"
	"github.com/fogleman/gg"
	"image"
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

// TODO: will be used for random image cards
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

func LoadImage(filename string) (image.Image, error) {
	ext := filepath.Ext(filename)
	switch strings.ToLower(ext) {
	case ".jpg", "jpeg":
		return gg.LoadJPG(filename)
	case ".png":
		return gg.LoadPNG(filename)
	default:
		return nil, errors.New("unsupported image format")
	}
}
