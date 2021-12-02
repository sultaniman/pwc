package gen

import (
	"github.com/imanhodjaev/pwc/canvas"
	"github.com/imanhodjaev/pwc/util"
	"image"
	"math/rand"
	"time"
)

const (
	Padding          = 10
	DefaultNumRows   = 3
	DefaultRowHeight = (canvas.Height - Padding) / DefaultNumRows
)

type Tile struct {
	Image    *image.Image
	Position *image.Point
}

type ImageCard struct {
	ImagePaths []string
	Tiles      []*Tile
	RowHeight  int
}

func (ic *ImageCard) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ic.ImagePaths), func(i, j int) {
		ic.ImagePaths[i], ic.ImagePaths[j] = ic.ImagePaths[j], ic.ImagePaths[i]
	})
}

func GenerateImageCard(collectionPath string) (*canvas.Canvas, *ImageCard, error) {
	imageCanvas := canvas.NewEmptyCanvas()
	paths, err := util.WalkAndReadDirectory(collectionPath)
	ic := &ImageCard{
		ImagePaths: paths,
	}

	if err != nil {
		return nil, nil, err
	}

	ic.Shuffle()
	return imageCanvas, ic, nil
}
