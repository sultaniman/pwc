package util

import (
	"github.com/go-fonts/dejavu/dejavusansmono"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
)

// LoadFontFace godoc
// Loads dejavusansmono and returns *font.Face with full hinting
func LoadFontFace(size int, dpi int) (*font.Face, error) {
	ttf, err := sfnt.Parse(dejavusansmono.TTF)
	if err != nil {
		return nil, err
	}

	// Create Font.Face from font
	face, err := opentype.NewFace(ttf, &opentype.FaceOptions{
		Size:    float64(size),
		DPI:     float64(dpi),
		Hinting: font.HintingFull,
	})

	if err != nil {
		return nil, err
	}

	return &face, nil
}
