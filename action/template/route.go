package template

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/middlewarex"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm/dialects/postgres"
)

// template model
type template struct {
	Title      string         `json:"title"`
	Slug       string         `json:"slug"`
	Spec       postgres.Jsonb `json:"spec"  swaggertype:"primitive,string"`
	Properties postgres.Jsonb `json:"properties"  swaggertype:"primitive,string"`
	MediumID   uint           `json:"medium_id"`
	SpaceID    uint           `json:"space_id"`
}

var userContext config.ContextKey = "template_user"

// Router - Group of template router
func Router() chi.Router {
	r := chi.NewRouter()

	app := "bindu"

	r.Get("/", list)
	r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Post("/", create)

	r.Route("/{template_id}", func(r chi.Router) {
		r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Put("/", update)
		r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Delete("/", delete)
		r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Get("/", details)
	})

	return r

}
