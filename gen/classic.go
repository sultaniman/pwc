package gen

import (
	"github.com/imanhodjaev/pwc/card"
	"github.com/imanhodjaev/pwc/util"
	"image"
	"strings"
)

type ClassicCard struct {
	Header     string
	Rows       []string
	Passphrase string
	Context    *AlphabetCollection
	Message    *util.Message
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
		Message: util.NewMessage(),
	}
}

// Generate godoc
// Generates randomized rows for a classic
// password card for each row we randomize
// alphabet and use it to generate a new row.
func (c *ClassicCard) Generate(alnumAndSymbols bool, digitsOnlyArea bool) error {
	chars := ""
	count := 0
	rows := 0
	c.Header = util.Shuffle(card.ClassicHeaderRow)

	passphrase, err := c.Message.RandomPassphrase()
	if err != nil {
		return err
	}

	c.Passphrase = passphrase

	for {
		if rows >= card.AlphabetBodyHeight {
			break
		}

		if count >= card.AlphabetWidth {
			c.Rows = append(c.Rows, chars)
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
		nextChar := c.Context.AlphaNumeric.Next()
		if alnumAndSymbols && count%2 == 0 {
			nextChar = c.Context.AlphaNumericAndSymbols.Next()
		}

		if digitsOnlyArea && rows > (card.AlphabetBodyHeight/2)-1 {
			nextChar = c.Context.Numeric.Next()
		}

		chars += nextChar
	}

	return nil
}

func (c *ClassicCard) GetBytes() []byte {
	rows := append([]string{c.Header}, c.Rows...)
	return []byte(strings.Join(rows, "\n"))
}

func GenerateClassicCard(withSymbols bool, digitsArea bool) (*card.Canvas, *ClassicCard, error) {
	canvas, err := card.NewCanvas()
	if err != nil {
		return nil, nil, err
	}

	passwordCard := NewClassicCard()
	err = passwordCard.Generate(withSymbols, digitsArea)
	if err != nil {
		return nil, nil, err
	}

	passwordCard.Message.Plaintext = string(passwordCard.GetBytes())

	_, height := canvas.Context.MeasureString(passwordCard.Header)
	canvas.ColorizeRows(height)
	canvas.RenderHeader(passwordCard.Header, height)
	canvas.Context.SetColor(image.Black)
	canvas.Context.SetFontFace(*canvas.FontFace)
	for i, row := range passwordCard.Rows {
		canvas.RenderRow(i, row, height)
	}

	canvas.RenderKey(passwordCard.Passphrase)
	return canvas, passwordCard, nil
}

func RestoreClassicCard(passphrase string, encryptedCard string) (*card.Canvas, *ClassicCard, error) {
	passwordCard := NewClassicCard()
	passwordCard.Passphrase = passphrase
	passwordCard.Message.Encrypted = encryptedCard

	decryptedCard, err := passwordCard.Message.Decrypt(passphrase)
	if err != nil {
		return nil, nil, err
	}

	passwordCard.Message.Plaintext = decryptedCard
	parts := strings.Split(decryptedCard, "\n")
	passwordCard.Header = parts[0]
	passwordCard.Rows = parts[1:]

	canvas, err := card.NewCanvas()
	if err != nil {
		return nil, nil, err
	}

	_, height := canvas.Context.MeasureString(passwordCard.Header)
	canvas.ColorizeRows(height)
	canvas.RenderHeader(passwordCard.Header, height)
	canvas.Context.SetColor(image.Black)
	canvas.Context.SetFontFace(*canvas.FontFace)
	for i, row := range passwordCard.Rows {
		canvas.RenderRow(i, row, height)
	}

	randomPassphrase, err := passwordCard.Message.RandomPassphrase()
	if err != nil {
		return nil, nil, err
	}

	passwordCard.Passphrase = randomPassphrase
	canvas.RenderKey(randomPassphrase)
	return canvas, passwordCard, nil
}
