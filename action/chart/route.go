package chart

import (
	"time"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
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
	IsPublic         bool           `json:"is_public"`
	FeaturedMedium   string         `json:"featured_medium"`
	FeaturedMediumID uint           `json:"featured_medium_id"`
	ThemeID          uint           `json:"theme_id"`
	TemplateID       string         `json:"template_id"`
	PublishedDate    *time.Time     `json:"published_date"`
	Mode             string         `json:"mode"`
	CategoryIDs      []uint         `json:"category_ids"`
	TagIDs           []uint         `json:"tag_ids"`
}

var userContext config.ContextKey = "chart_user"

// Router - Group of chart router
func Router() chi.Router {
	r := chi.NewRouter()

	entity := "charts"

	r.With(util.CheckKetoPolicy(entity, "get")).Get("/", list)
	r.With(util.CheckKetoPolicy(entity, "create")).Post("/", create)

	r.Route("/{chart_id}", func(r chi.Router) {
		r.With(util.CheckKetoPolicy(entity, "get")).Get("/", details)
		r.With(util.CheckKetoPolicy(entity, "update")).Put("/", update)
		r.With(util.CheckKetoPolicy(entity, "delete")).Delete("/", delete)
	})

	return r

}
