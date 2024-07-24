package commands

import (
	_ "fmt"
	"s3_file_uploader/migrations"
	"s3_file_uploader/utils"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migrations",
	Run: func(cmd *cobra.Command, args []string) {
		utils.InitMongoDB()

		err := migrations.Up()
		if err != nil {
			utils.MongoLogError("Migration failed: ", err)
			return
		}

		utils.MongoLogInfo("Migration completed successfully")
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
