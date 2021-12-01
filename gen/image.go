package gen

import (
	"github.com/imanhodjaev/pwc/canvas"
	"github.com/imanhodjaev/pwc/util"
	"image"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

type Tile struct {
	Image    *image.Image
	Position *image.Point
}

type ImageCard struct {
	ImagePaths []string
	Tiles      []*Tile
}

func (ic *ImageCard) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ic.ImagePaths), func(i, j int) {
		ic.ImagePaths[i], ic.ImagePaths[j] = ic.ImagePaths[j], ic.ImagePaths[i]
	})
}

func GenerateImageCard(collectionPath string) (*canvas.Canvas, *ImageCard, error) {
	imageCanvas := canvas.NewEmptyCanvas()
	ic := &ImageCard{}
	err := filepath.Walk(collectionPath, func(currentPath string, info os.FileInfo, err error) error {
		if mode := info.Mode(); mode.IsRegular() {
			if util.AllowedFormats(filepath.Ext(currentPath)) {
				ic.ImagePaths = append(ic.ImagePaths, currentPath)
			}
		}

		return err
	})

	if err != nil {
		return nil, nil, err
	}

	ic.Shuffle()
	return imageCanvas, ic, nil
}
