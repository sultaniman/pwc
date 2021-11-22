package gen

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/imanhodjaev/pwc/card"
	"github.com/imanhodjaev/pwc/util"
	"golang.org/x/crypto/scrypt"
	"image"
	"strings"
)

const AESKeyLength = 32

type ClassicCard struct {
	Header  string
	Rows    []string
	AESKey  string
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
func (c *ClassicCard) Generate(alnumAndSymbols bool, digitsOnlyArea bool) error {
	chars := ""
	count := 0
	rows := 0
	c.Header = util.Shuffle(card.ClassicHeaderRow)

	keyBytes := make([]byte, AESKeyLength)
	_, err := rand.Read(keyBytes)
	if err != nil {
		return err
	}

	c.AESKey = hex.EncodeToString(keyBytes)[:16]

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

func (c *ClassicCard) Encrypt() (string, error) {
	aesKey, _ := hex.DecodeString(c.AESKey)
	key, salt, err := DeriveKey(aesKey, nil)
	if err != nil {
		return "", err
	}

	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	encrypted := aesGCM.Seal(nonce, nonce, c.GetBytes(), nil)
	encrypted = append(encrypted, salt...)
	return hex.EncodeToString(encrypted), nil
}

func (c *ClassicCard) GetBytes() []byte {
	rows := append([]string{c.Header}, c.Rows...)
	return []byte(strings.Join(rows, "\n"))
}

// DeriveKey TODO: parameterize scrypt cost params
func DeriveKey(key, salt []byte) ([]byte, []byte, error) {
	if salt == nil {
		salt = make([]byte, 32)
		if _, err := rand.Read(salt); err != nil {
			return nil, nil, err
		}
	}

	key, err := scrypt.Key(key, salt, 10000, 8, 1, 32)
	if err != nil {
		return nil, nil, err
	}

	return key, salt, nil
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

	_, height := canvas.Context.MeasureString(passwordCard.Header)
	canvas.ColorizeRows(height)
	canvas.RenderHeader(passwordCard.Header, height)

	canvas.Context.SetColor(image.Black)
	canvas.Context.SetFontFace(*canvas.FontFace)
	for i, row := range passwordCard.Rows {
		canvas.RenderRow(i, row, height)
	}

	canvas.RenderKey(passwordCard.AESKey)
	return canvas, passwordCard, nil
}
