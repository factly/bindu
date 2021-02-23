package cmd

import (
	"log"
	"net/http"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util/minio"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts server for bindu-server.",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.CreateSuperOrganisation()
		if err != nil {
			log.Println(err)
		}

		minio.SetupClient()

		// db setup
		config.SetupDB()
		// register routes
		r := action.RegisterRoutes()

		err = http.ListenAndServe(":8000", r)

		if err != nil {
			log.Fatal(err)
		}
	},
}
