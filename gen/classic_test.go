package gen

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sultaniman/pwc/canvas"
)

func TestClassicCard_Generate(t *testing.T) {
	classicCard := NewClassicCard()
	err := classicCard.Generate(false, true)

	if assert.NoError(t, err) {
		assert.Equal(t, 8, len(classicCard.Rows))
		for _, row := range classicCard.Rows {
			assert.Equal(t, canvas.AlphabetWidth, len(row))
		}
	}
}
