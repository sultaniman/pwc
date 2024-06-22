package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sultaniman/pwc/gen"
	"github.com/sultaniman/pwc/util"
)

var passphrase string

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore password card from encrypted backup",
	Long:  "Restore password card from encrypted backup",
	RunE: func(cmd *cobra.Command, args []string) error {
		if outputFile == "" {
			fmt.Println("Please specify output filename for example restored.jpg")
			os.Exit(1)
		}

		if encryptedFile == "" {
			fmt.Println("Please provide encrypted backup file")
			os.Exit(1)
		}

		if !util.FileExists(encryptedFile) {
			fmt.Println("Encrypted backup file does not exist")
			os.Exit(1)
		}

		if passphrase == "" {
			fmt.Println("Please specify passphrase")
			os.Exit(1)
		}

		encryptedBytes, err := os.ReadFile(encryptedFile)
		if err != nil {
			return err
		}

		canvas, passwordCard, err := gen.RestoreClassicCard(passphrase, string(encryptedBytes))
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

		return util.SaveImage(canvas.Context, outputFile)
	},
}

func init() {
	restoreCmd.PersistentFlags().StringVarP(
		&outputFile,
		"output",
		"o",
		"card.jpg",
		"Output file (supported formats PNG, JPG)",
	)

	restoreCmd.PersistentFlags().BoolVarP(
		&printPassphrase,
		"print-passphrase",
		"x",
		false,
		"Prints passphrase in the console",
	)

	restoreCmd.PersistentFlags().StringVarP(
		&passphrase,
		"passphrase",
		"p",
		"",
		"Passphrase to decrypt contents",
	)

	restoreCmd.PersistentFlags().StringVarP(
		&encryptedFile,
		"encrypted",
		"e",
		"",
		"Path to encrypted backup file",
	)
}
