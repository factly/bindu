package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/factly/bindu-server/action"
	"github.com/factly/bindu-server/action/chart"
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/factly/bindu-server/util/minio"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts server for bindu-server.",
	Run: func(cmd *cobra.Command, args []string) {
		minio.SetupClient()

		// db setup
		config.SetupDB()
		// register routes
		r := action.RegisterRoutes()

		go ServeCharts()

		err := http.ListenAndServe(":8000", r)
		if err != nil {
			log.Fatal(err)
		}
	},
}

// ServeCharts server for chart routes
func ServeCharts() {
	util.SetupTemplates()

	r := chi.NewRouter()

	r.Use(cors.AllowAll().Handler)
	r.Get("/charts/visualization/{chart_id}", chart.Visualize)

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "web/"))

	FileServer(r, "/", filesDir)

	err := http.ListenAndServe(":8001", r)
	if err != nil {
		log.Fatal(err)
	}
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		if strings.HasSuffix(r.URL.Path, ".js") {
			w.Header().Add("Content-Type", "application/javascript")
		}
		if strings.HasSuffix(r.URL.Path, ".ts") {
			w.Header().Add("Content-Type", "application/typescript")
		}
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
