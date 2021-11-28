package cmd

import (
	"fmt"
	"github.com/imanhodjaev/pwc/gen"
	"github.com/imanhodjaev/pwc/util"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
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

		encryptedBytes, err := ioutil.ReadFile(encryptedFile)
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

		return canvas.Save(outputFile)

		return nil
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
