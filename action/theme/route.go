package theme

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
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

	entity := "themes"

	r.With(util.CheckKetoPolicy(entity, "get")).Get("/", list)
	r.With(util.CheckKetoPolicy(entity, "create")).Post("/", create)

	r.Route("/{theme_id}", func(r chi.Router) {
		r.With(util.CheckKetoPolicy(entity, "get")).Get("/", details)
		r.With(util.CheckKetoPolicy(entity, "update")).Put("/", update)
		r.With(util.CheckKetoPolicy(entity, "delete")).Delete("/", delete)
	})

	return r

}
