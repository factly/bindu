package category

import (
	"github.com/factly/bindu-server/config"
	"github.com/go-chi/chi"
)

// category request model
type category struct {
	Name        string `json:"name" validate:"required,min=3,max=50"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

var userContext config.ContextKey = "category_user"

// Router - Group of category router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)
	r.Post("/", create)

	r.Route("/{category_id}", func(r chi.Router) {
		r.Get("/", details)
		r.Put("/", update)
		r.Delete("/", delete)
	})

	return r

}
