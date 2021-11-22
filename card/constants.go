package card

import (
	"image/color"
)

const (
	AlphabetWidth          = 29
	AlphabetBodyHeight     = 8
	ClassicHeaderRow       = "■□▲△○●★☂☀☁☹☺♠♣♥♦♫€¥£$!?¡¿⊙◐◩�"
	Numbers                = "0123456789"
	AlphabetLower          = "abcdefghijklmnopqrstuvwxyz"
	AlphabetUpper          = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Symbols                = "@#$%&*<>?€+{}[]()/\\"
	AlphaNumeric           = Numbers + AlphabetLower + AlphabetUpper
	AlphaNumericAndSymbols = AlphaNumeric + Symbols
)

var (
	ColorWhite = color.RGBA{
		R: 255,
		G: 255,
		B: 255,
	}

	ColorGray = color.RGBA{
		R: 192,
		G: 192,
		B: 192,
	}

	ColorRed = color.RGBA{
		R: 255,
		G: 192,
		B: 192,
	}

	ColorGreen = color.RGBA{
		R: 192,
		G: 255,
		B: 192,
	}

	ColorYellow = color.RGBA{
		R: 255,
		G: 255,
		B: 192,
	}

	ColorBlue = color.RGBA{
		R: 192,
		G: 192,
		B: 255,
	}

	ColorMagenta = color.RGBA{
		R: 255,
		G: 192,
		B: 255,
	}

	ColorCyan = color.RGBA{
		R: 192,
		G: 255,
		B: 255,
	}

	Colors = []color.RGBA{
		ColorWhite,
		ColorGray,
		ColorRed,
		ColorGreen,
		ColorYellow,
		ColorBlue,
		ColorMagenta,
		ColorCyan,
	}
)
