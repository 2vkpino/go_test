package commands

import (
	"fmt"
	"s3_file_uploader/migrations"
	"s3_file_uploader/utils"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		err := migrations.Up()
		if err != nil {
			utils.LogError(fmt.Sprintf("Migration failed: %v", err))
			return
		}
		utils.LogInfo("Migration completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
