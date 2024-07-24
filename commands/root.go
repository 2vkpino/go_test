package commands

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "s3_uploader",
	Short: "S3 File Uploader Service",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Здесь можно инициализировать другие команды
}
