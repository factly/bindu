package theme

import (
	"github.com/factly/bindu-server/config"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// theme request model
type theme struct {
	Name        string         `json:"name" validate:"required,min=3,max=50"`
	Config      postgres.Jsonb `json:"config" swaggertype:"primitive,string"`
	Description string         `json:"description"`
}

var userContext config.ContextKey = "theme_user"

// Router - Group of theme router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)
	r.Post("/", create)

	r.Route("/{theme_id}", func(r chi.Router) {
		r.Get("/", details)
		r.Put("/", update)
		r.Delete("/", delete)
	})

	return r

}
