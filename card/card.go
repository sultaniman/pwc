package card

import (
	"crypto/rand"
	"golang.org/x/image/font"
	"image"
	"image/jpeg"
	"os"
)

const (
	Width          = 1050
	Height         = 600
	AESKeyLength   = 16
	DefaultQuality = 80
)

type Card struct {
	Image    *image.RGBA
	FontFace *font.Face
	AESKey   []byte
}

func (c *Card) RenderHeader() error {
	return nil
}

func (c *Card) RenderRow(index int, contents string) error {
	return nil
}

func (c *Card) Save(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	err = jpeg.Encode(f, c.Image, &jpeg.Options{Quality: DefaultQuality})
	if err != nil {
		return err
	}

	return nil
}

func NewBlankCard() (*Card, error) {
	key := make([]byte, AESKeyLength)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return &Card{
		Image:  image.NewRGBA(image.Rect(0, 0, Width, Height)),
		AESKey: nil,
	}, nil
}
