package category

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/go-chi/chi"
)

// category request model
type category struct {
	Name          string `json:"name" validate:"required,min=3,max=50"`
	Slug          string `json:"slug"`
	IsForTemplate bool   `json:"is_for_template"`
	Description   string `json:"description"`
}

var userContext config.ContextKey = "category_user"

// Router - Group of category router
func Router() chi.Router {
	r := chi.NewRouter()

	entity := "categories"

	r.With(util.CheckKetoPolicy(entity, "get")).Get("/", list)
	r.With(util.CheckKetoPolicy(entity, "create")).Post("/", create)

	r.Route("/{category_id}", func(r chi.Router) {
		r.With(util.CheckKetoPolicy(entity, "get")).Get("/", details)
		r.With(util.CheckKetoPolicy(entity, "update")).Put("/", update)
		r.With(util.CheckKetoPolicy(entity, "delete")).Delete("/", delete)
	})

	return r

}
