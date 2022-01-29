package gen

import (
	"github.com/sultaniman/pwc/util"
)

type Alphabet struct {
	Letters string
	Runes   []rune
	Rand    *util.RandRange
}

// Next godoc
// Select next letter and shuffle alphabet
func (a *Alphabet) Next() string {
	char := string(a.Runes[a.Rand.Next()])
	a.Shuffle()
	return char
}

// Shuffle godoc
// Shuffle all letters and update runes
func (a *Alphabet) Shuffle() {
	a.Letters = util.Shuffle(a.Letters)
	a.Runes = []rune(a.Letters)
}

// NewAlphabet godoc
// Shuffle and initialize given Alphabet
// Also initialize random selector
func NewAlphabet(alphabet string) *Alphabet {
	newAlphabet := util.Shuffle(alphabet)
	runes := []rune(newAlphabet)
	return &Alphabet{
		Letters: newAlphabet,
		Runes:   runes,
		Rand:    util.NewRandRange(0, len(runes)-1),
	}
}
