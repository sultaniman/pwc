package util

import (
	"math/rand"
	"time"
)

func Shuffle(inStr string) string {
	rand.Seed(time.Now().Unix())
	inRune := []rune(inStr)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
