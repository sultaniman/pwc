package util

import (
	"github.com/fogleman/gg"
	"strings"
)

var supportedExtensions = map[string]bool{
	"png":  true,
	"jpg":  true,
	"jpeg": true,
}

const DefaultQuality = 80

func SaveImage(ctx *gg.Context, path string) error {
	parts := strings.Split(path, ".")
	ext := strings.ToLower(parts[len(parts)-1])
	if ext == "" || !supportedExtensions[ext] {
		ext = "png"
		path += "." + ext
	}

	if ext == "png" {
		return ctx.SavePNG(path)
	}

	if ext == "jpg" || ext == "jpeg" {
		return ctx.SaveJPG(path, DefaultQuality)
	}

	return nil
}
