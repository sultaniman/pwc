package cmd

import (
	"github.com/imanhodjaev/pwc/card"
	"github.com/imanhodjaev/pwc/gen"
	"github.com/imanhodjaev/pwc/util"
	"github.com/spf13/cobra"
	"image"
)

var (
	outputFile  = "card.png"
	withSymbols = false
	digitsArea  = false
)

var classicCmd = &cobra.Command{
	Use:   "classic",
	Short: "Generate classic password card",
	Long:  "Generate classic password card",
	RunE: func(cmd *cobra.Command, args []string) error {
		canvas, err := card.NewCanvas()
		if err != nil {
			return err
		}

		shuffledHeaderRow := util.Shuffle(card.ClassicHeaderRow)
		_, height := canvas.Context.MeasureString(shuffledHeaderRow)
		canvas.ColorizeRows(height)
		canvas.RenderHeader(shuffledHeaderRow, height)
		passwordCard := gen.NewClassicCard()
		passwordCard.Generate(withSymbols, digitsArea)
		canvas.Context.SetColor(image.Black)
		canvas.Context.SetFontFace(*canvas.FontFace)
		for i, row := range passwordCard.Rows {
			canvas.RenderRow(i, row, height)
		}

		canvas.RenderKey(canvas.AESKey)
		return canvas.Save(outputFile)
	},
}

func init() {
	classicCmd.PersistentFlags().StringVarP(
		&outputFile,
		"output",
		"o",
		"card.jpg",
		"Output file",
	)

	classicCmd.PersistentFlags().BoolVarP(
		&withSymbols,
		"include-symbols",
		"s",
		false,
		"With regular a-zA-Z include @#$%&*<>?€+{}[]()/\\",
	)

	classicCmd.PersistentFlags().BoolVarP(
		&digitsArea,
		"include-digits",
		"d",
		false,
		"Rows 5-8 will be digits only",
	)
}
