package cmd

import (
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Generate password card using images",
	Long:  "Generate password card using images",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
