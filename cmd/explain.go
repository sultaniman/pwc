package cmd

import (
	"fmt"
	"github.com/fatih/color"
	goc "github.com/gookit/color"
	"github.com/imanhodjaev/pwc/card"
	"github.com/spf13/cobra"
)

const algorithmDescription = `
1. Render the first row with shuffled "■□▲△○●★☂☀☁☹☺♠♣♥♦♫€¥£$!?¡¿⊙◐◩�",
2. Iterate over the rest of the rows
    a. If the card should include symbols then
       use the alphanumeric and symbols alphabet for every even character
       and use the alphanumeric alphabet for the odd columns,
    b. If the card should have a digits area then
       use the numeric alphabet for the lower half of the rows 5-8.
`

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Describe algorithm and show alphabets",
	Long:  "Describe algorithm and show alphabets",
	Run: func(cmd *cobra.Command, args []string) {
		headerColor := color.New(color.FgGreen, color.Bold).Add(color.Underline)
		alphabetColor := color.New(color.FgYellow, color.Bold)
		_, _ = headerColor.Print("Alphabet\n\n")
		fmt.Printf("%30s %s\n", "Header symbols:", alphabetColor.Sprintf(card.ClassicHeaderRow))
		fmt.Printf("%30s %s\n", "Numbers:", alphabetColor.Sprintf(card.Numbers))
		fmt.Printf("%30s %s\n", "Alphanumeric:", alphabetColor.Sprintf(card.AlphaNumeric))
		fmt.Printf("%30s %s\n", "Alphanumeric and symbols:", alphabetColor.Sprintf(card.AlphaNumericAndSymbols))
		_, _ = headerColor.Print("\nAlgorithm")
		fmt.Printf("\n%s", algorithmDescription)

		_, _ = headerColor.Print("\nRow colors\n\n")
		for index, rgb := range card.Colors {
			comma := ""
			if index < len(card.Colors)-1 {
				comma = ", "
			}
			goc.RGB(rgb.R, rgb.G, rgb.B).Printf("%s #%x%x%x%s", card.ColorNames[index], rgb.R, rgb.G, rgb.B, comma)
		}

		fmt.Printf("\n\n%50s✨ 🚀 ✨\n", " ")
	},
}
