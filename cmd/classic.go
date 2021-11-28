package cmd

import (
	"fmt"
	"github.com/imanhodjaev/pwc/gen"
	"github.com/spf13/cobra"
	"os"
	"syscall"
)

var (
	outputFile      = "card.png"
	encryptedFile   = "card.aes"
	printPassphrase = false
	withSymbols     = false
	digitsArea      = false
)

var classicCmd = &cobra.Command{
	Use:   "classic",
	Short: "Generate classic password card",
	Long:  "Generate classic password card",
	RunE: func(cmd *cobra.Command, args []string) error {
		if outputFile == "" {
			fmt.Println("Please provide output file using -o option")
			syscall.Exit(1)
		}

		canvas, passwordCard, err := gen.GenerateClassicCard(withSymbols, digitsArea)
		if err != nil {
			return err
		}

		if encryptedFile != "" {
			encrypted, err := passwordCard.Message.Encrypt(passwordCard.Passphrase)
			if err != nil {
				return err
			}

			fp, err := os.Create(encryptedFile)
			if err != nil {
				return err
			}

			_, err = fp.WriteString(encrypted)
			if err != nil {
				return err
			}

			err = fp.Close()
			if err != nil {
				return err
			}
		}

		if printPassphrase {
			fmt.Printf("Passphrase: %s\n", passwordCard.Passphrase)
		}

		return canvas.Save(outputFile)
	},
}

func init() {
	classicCmd.PersistentFlags().StringVarP(
		&outputFile,
		"output",
		"o",
		"card.jpg",
		"Output file (supported formats PNG, JPG)",
	)

	classicCmd.PersistentFlags().StringVarP(
		&encryptedFile,
		"encrypted",
		"e",
		"card.aes",
		"When given will encrypt generated card and write to file",
	)

	classicCmd.PersistentFlags().BoolVarP(
		&printPassphrase,
		"print-passphrase",
		"x",
		false,
		"Prints passphrase in the console",
	)

	classicCmd.PersistentFlags().BoolVarP(
		&withSymbols,
		"include-symbols",
		"s",
		false,
		"With regular a-zA-Z include @#$%&*<>?€+{}[]()/\\",
	)

	classicCmd.PersistentFlags().BoolVarP(
		&digitsArea,
		"include-digits",
		"d",
		false,
		"Rows 5-8 will be digits only",
	)
}
