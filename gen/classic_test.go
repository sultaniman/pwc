package gen

import (
	"github.com/imanhodjaev/pwc/card"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClassicCard_Generate(t *testing.T) {
	classicCard := NewClassicCard()
	classicCard.Generate(false, true)

	assert.Equal(t, 8, len(classicCard.Rows))
	for _, row := range classicCard.Rows {
		assert.Equal(t, card.AlphabetWidth, len(row))
	}
}
