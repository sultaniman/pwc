package canvas

import (
	"github.com/fogleman/gg"
	"github.com/imanhodjaev/pwc/util"
	"golang.org/x/image/font"
	"image"
	"strconv"
)

const (
	Width               = 1080
	Height              = 680
	MarginTop           = 10
	MarginRight         = 70
	MarginLeft          = 22
	ColorRowMarginLeft  = 38
	PaddingTop          = 70
	PaddingLeft         = 66
	HeaderMarginLeft    = MarginLeft * 4
	RowIndexMargin      = 5
	FontSize            = 10
	PassKeyBottomMargin = Height - 35
)

type Canvas struct {
	Context       *gg.Context
	FontFace      *font.Face
	IndexFontFace *font.Face
}

func (c *Canvas) RenderHeader(headerLetters string, offset float64) {
	c.Context.SetColor(image.Black)
	c.Context.DrawString(headerLetters, HeaderMarginLeft, offset)
}

func (c *Canvas) ColorizeRows(rowHeight float64) {
	c.Context.SetFontFace(*c.IndexFontFace)
	for i, col := range Colors {
		c.Context.SetColor(col)
		c.Context.DrawRectangle(
			ColorRowMarginLeft,
			MarginTop+rowHeight*float64(i+1),
			float64(c.Context.Width()-MarginRight),
			rowHeight,
		)

		c.Context.Fill()

		// TODO: extract into separate method
		c.Context.SetColor(image.Black)
		c.Context.DrawString(strconv.Itoa(i+1), ColorRowMarginLeft+RowIndexMargin, 40+rowHeight*float64(i+1))
	}

	c.Context.SetFontFace(*c.FontFace)
}

func (c *Canvas) RenderRow(index int, row string, rowHeight float64) {
	c.Context.DrawString(row, PaddingLeft+MarginLeft, PaddingTop+rowHeight*float64(index+1))
}

func (c *Canvas) RenderKey(key string) {
	c.Context.SetColor(image.Black)
	c.Context.SetFontFace(*c.IndexFontFace)
	width, _ := c.Context.MeasureString(key)
	x := float64(Width/2) - width/2
	c.Context.DrawString(key, x, PassKeyBottomMargin)
	c.Context.SetFontFace(*c.FontFace)
}

func NewCanvas() (*Canvas, error) {
	dc := gg.NewContext(Width, Height)
	dc.DrawRectangle(0, 0, Width, Height)
	dc.SetColor(image.White)
	dc.Fill()

	fontFace, err := util.LoadFontFace(FontSize, 400)
	if err != nil {
		return nil, err
	}

	indexFontFace, err := util.LoadFontFace(FontSize, 200)
	if err != nil {
		return nil, err
	}

	card := Canvas{
		Context:       dc,
		FontFace:      fontFace,
		IndexFontFace: indexFontFace,
	}

	dc.SetFontFace(*card.FontFace)

	return &card, nil
}
