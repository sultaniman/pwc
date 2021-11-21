package cmd

import (
	"github.com/spf13/cobra"
)

var imageFolder string

var rootCmd = &cobra.Command{
	Use:   "cmd",
	Short: "Password card generator",
	Long: "Password card generator",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&imageFolder,
		"--folder",
		"-f",
		"",
		"Folder with images",
	)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
