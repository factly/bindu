package medium

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// medium model
type medium struct {
	Name        string         `json:"name" validate:"required"`
	Slug        string         `json:"slug"`
	Type        string         `json:"type" `
	Description string         `json:"description"`
	Caption     string         `json:"caption"`
	AltText     string         `json:"alt_text"`
	FileSize    int64          `json:"file_size" `
	URL         postgres.Jsonb `json:"url" swaggertype:"primitive,string"`
	Dimensions  string         `json:"dimensions" `
}

var userContext config.ContextKey = "medium_user"

// Router - Group of medium router
func Router() chi.Router {
	r := chi.NewRouter()

	entity := "media"

	r.With(util.CheckKetoPolicy(entity, "get")).Get("/", list)
	r.With(util.CheckKetoPolicy(entity, "create")).Post("/", create)

	r.Route("/{medium_id}", func(r chi.Router) {
		r.With(util.CheckKetoPolicy(entity, "get")).Get("/", details)
		r.With(util.CheckKetoPolicy(entity, "update")).Put("/", update)
		r.With(util.CheckKetoPolicy(entity, "delete")).Delete("/", delete)
	})

	return r

}
