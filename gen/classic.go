package gen

import (
	"github.com/imanhodjaev/pwc/card"
)

type ClassicCard struct {
	Rows    []string
	Context *AlphabetCollection
}

type AlphabetCollection struct {
	Numeric                *Alphabet
	AlphaNumeric           *Alphabet
	AlphaNumericAndSymbols *Alphabet
}

func NewClassicCard() *ClassicCard {
	return &ClassicCard{
		Context: &AlphabetCollection{
			Numeric:                NewAlphabet(card.Numbers),
			AlphaNumeric:           NewAlphabet(card.AlphaNumeric),
			AlphaNumericAndSymbols: NewAlphabet(card.AlphaNumericAndSymbols),
		},
	}
}

// Generate godoc
// Generates randomized rows for a classic
// password card for each row we randomize
// alphabet and use it to generate a new row.
func (sr *ClassicCard) Generate(alnumAndSymbols bool, digitsOnlyArea bool) {
	chars := ""
	count := 0
	rows := 0

	for {
		if rows >= card.AlphabetBodyHeight {
			break
		}

		if count >= card.AlphabetWidth {
			sr.Rows = append(sr.Rows, chars)
			count = 0
			chars = ""
			rows++
			continue
		} else {
			count++
		}

		// Get next character from alphanumeric alphabet anyway.
		// If mode is alphanumeric and symbols and counter is even then
		// take next character from alphanumeric and symbols alphabet.
		nextChar := sr.Context.AlphaNumeric.Next()
		if alnumAndSymbols && count%2 == 0 {
			nextChar = sr.Context.AlphaNumericAndSymbols.Next()
		}

		if digitsOnlyArea && rows > (card.AlphabetBodyHeight/2)-1 {
			nextChar = sr.Context.Numeric.Next()
		}

		chars += nextChar
	}
}
