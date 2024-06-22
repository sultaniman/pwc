package util

import "math/rand"

func Shuffle(inStr string) string {
	inRune := []rune(inStr)
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}
