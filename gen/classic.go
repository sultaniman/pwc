package gen

import (
	"github.com/imanhodjaev/pwc/canvas"
	"github.com/imanhodjaev/pwc/crypto"
	"github.com/imanhodjaev/pwc/util"
	"image"
	"strings"
)

type ClassicCard struct {
	Header     string
	Rows       []string
	Passphrase string
	Context    *AlphabetCollection
	Message    *crypto.Message
}

type AlphabetCollection struct {
	Numeric                *Alphabet
	AlphaNumeric           *Alphabet
	AlphaNumericAndSymbols *Alphabet
}

func NewClassicCard() *ClassicCard {
	return &ClassicCard{
		Context: &AlphabetCollection{
			Numeric:                NewAlphabet(canvas.Numbers),
			AlphaNumeric:           NewAlphabet(canvas.AlphaNumeric),
			AlphaNumericAndSymbols: NewAlphabet(canvas.AlphaNumericAndSymbols),
		},
		Message: crypto.NewMessage("", ""),
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
	c.Header = util.Shuffle(canvas.ClassicHeaderRow)

	passphrase, err := c.Message.RandomPassphrase()
	if err != nil {
		return err
	}

	c.Passphrase = passphrase

	for {
		if rows >= canvas.AlphabetBodyHeight {
			break
		}

		if count >= canvas.AlphabetWidth {
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

		if digitsOnlyArea && rows > (canvas.AlphabetBodyHeight/2)-1 {
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

func GenerateClassicCard(withSymbols bool, digitsArea bool, renderImage bool) (*canvas.Canvas, *ClassicCard, error) {
	var cardCanvas *canvas.Canvas
	passwordCard := NewClassicCard()
	err := passwordCard.Generate(withSymbols, digitsArea)
	if err != nil {
		return nil, nil, err
	}

	passwordCard.Message.Plaintext = string(passwordCard.GetBytes())

	if renderImage {
		cardCanvas, err = canvas.NewCanvas()
		if err != nil {
			return nil, nil, err
		}

		_, height := cardCanvas.Context.MeasureString(passwordCard.Header)
		cardCanvas.ColorizeRows(height)
		cardCanvas.RenderHeader(passwordCard.Header, height)
		cardCanvas.Context.SetColor(image.Black)
		cardCanvas.Context.SetFontFace(*cardCanvas.FontFace)
		for i, row := range passwordCard.Rows {
			cardCanvas.RenderRow(i, row, height)
		}

		cardCanvas.RenderKey(passwordCard.Passphrase)
	}
	return cardCanvas, passwordCard, nil
}

func RestoreClassicCard(passphrase string, encryptedCard string) (*canvas.Canvas, *ClassicCard, error) {
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

	canva, err := canvas.NewCanvas()
	if err != nil {
		return nil, nil, err
	}

	_, height := canva.Context.MeasureString(passwordCard.Header)
	canva.ColorizeRows(height)
	canva.RenderHeader(passwordCard.Header, height)
	canva.Context.SetColor(image.Black)
	canva.Context.SetFontFace(*canva.FontFace)
	for i, row := range passwordCard.Rows {
		canva.RenderRow(i, row, height)
	}

	randomPassphrase, err := passwordCard.Message.RandomPassphrase()
	if err != nil {
		return nil, nil, err
	}

	passwordCard.Passphrase = randomPassphrase
	canva.RenderKey(randomPassphrase)
	return canva, passwordCard, nil
}
