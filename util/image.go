package util

import (
	"path/filepath"
	"strings"

	"github.com/fogleman/gg"
)

var supportedExtensions = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
}

const DefaultQuality = 80

func SaveImage(ctx *gg.Context, path string) error {
	ext := strings.ToLower(filepath.Ext(path))
	if ext == "" || !supportedExtensions[ext] {
		ext = ".png"
		path += "." + ext
	}

	if ext == ".png" {
		return ctx.SavePNG(path)
	}

	if ext == ".jpg" || ext == ".jpeg" {
		return ctx.SaveJPG(path, DefaultQuality)
	}

	return nil
}
