package cmd

import (
	"fmt"
	"github.com/imanhodjaev/pwc/card"
	"github.com/spf13/cobra"
)

var explainCmd = &cobra.Command{
	Use:   "explain",
	Short: "Describe algorithm and show alphabets",
	Long:  "Describe algorithm and show alphabets",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%30s %s\n", "Header symbols:", card.ClassicHeaderRow)
		fmt.Printf("%30s %s\n", "Numbers:", card.Numbers)
		fmt.Printf("%30s %s\n", "Symbols:", card.Symbols)
		fmt.Printf("%30s %s\n", "Alphanumeric:", card.AlphaNumeric)
		fmt.Printf("%30s %s\n", "Alphanumeric and symbols:", card.AlphaNumericAndSymbols)
	},
}
