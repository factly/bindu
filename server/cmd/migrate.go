package cmd

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/model"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCmd)
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply database migrations for bindu-server.",
	Run: func(cmd *cobra.Command, args []string) {
		// db setup
		config.SetupDB()

		// apply migrations
		model.Migration()
	},
}
