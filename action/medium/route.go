package medium

import (
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// medium model
type medium struct {
	Name string         `json:"name" validate:"required,min=3,max=50"`
	Slug string         `json:"slug"`
	Type string         `json:"type"`
	URL  postgres.Jsonb `json:"url" swaggertype:"primitive,string"`
}

// Router - Group of medium router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)
	r.Post("/", create)

	r.Route("/{medium_id}", func(r chi.Router) {
		r.Get("/", details)
		r.Put("/", update)
		r.Delete("/", delete)
	})

	return r

}
