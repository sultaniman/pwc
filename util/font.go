package util

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
)

func LoadGomonoFace(size int, dpi int) (*font.Face, error) {
	ttf, err := truetype.Parse(gomono.TTF)
	if err != nil {
		return nil, err
	}

	//Create Font.Face from font
	face := truetype.NewFace(ttf, &truetype.Options{
		Size:    float64(size),
		DPI:     float64(dpi),
		Hinting: font.HintingNone,
	})

	return &face, nil
}
