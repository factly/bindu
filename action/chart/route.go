package chart

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// chart request model
type chart struct {
	Title            string         `json:"title" validate:"required,min=3,max=50"`
	Slug             string         `json:"slug"`
	DataURL          string         `json:"data_url"`
	Config           postgres.Jsonb `json:"config" swaggertype:"primitive,string"`
	Description      postgres.Jsonb `json:"description" swaggertype:"primitive,string"`
	Status           string         `json:"status"`
	FeaturedMedium   string         `json:"featured_medium"`
	FeaturedMediumID uint           `json:"featured_medium_id"`
	ThemeID          uint           `json:"theme_id"`
	PublishedDate    time.Time      `json:"published_date"`
	OrganisationID   uint           `json:"organisation_id"`
	CategoryIDs      []uint         `json:"category_ids"`
	TagIDs           []uint         `json:"tag_ids"`
}

// Router - Group of chart router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)
	r.Post("/", create)

	r.Route("/{chart_id}", func(r chi.Router) {
		r.Get("/", details)
		r.Put("/", update)
		r.Delete("/", delete)
	})

	return r

}
