package cmd

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/imanhodjaev/pwc/gen"
	"github.com/imanhodjaev/pwc/util"
	"github.com/spf13/cobra"
	"os"
)

var collectionPath string

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Generate password card using images",
	Long:  "Generate password card using images",
	RunE: func(cmd *cobra.Command, args []string) error {
		if collectionPath == "" {
			fmt.Println("Please specify path to image collection")
			os.Exit(1)
		}

		if !util.FileExists(collectionPath) {
			fmt.Println("Specified path to image collection does not exist")
			os.Exit(1)
		}

		canvas, ic, err := gen.GenerateImageCard(collectionPath)
		if err != nil {
			return err
		}
		_ = canvas.Context
		spew.Dump(ic.ImagePaths)

		return nil
	},
}

func init() {
	imageCmd.PersistentFlags().StringVarP(
		&outputFile,
		"output",
		"o",
		"card.jpg",
		"Output file (supported formats PNG, JPG)",
	)

	imageCmd.PersistentFlags().StringVarP(
		&collectionPath,
		"collection",
		"c",
		"",
		"Path to folder with images",
	)
}
