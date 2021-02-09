package organisation

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/middlewarex"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type organisationPermission struct {
	OrganisationID uint  `json:"organisation_id" validate:"required"`
	Spaces         int64 `json:"spaces"`
}

type organisationPermissionRequest struct {
	Title          string         `json:"title" validate:"required"`
	OrganisationID uint           `json:"organisation_id" validate:"required"`
	Description    postgres.Jsonb `json:"description" swaggertype:"primitive,string"`
	Spaces         int64          `json:"spaces"`
}

var userContext config.ContextKey = "org_perm_user"
var requestUserContext config.ContextKey = "request_user"

// Router - Group of medium router
func Router() chi.Router {
	r := chi.NewRouter()

	app := "bindu"

	r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Get("/", list)
	r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Post("/", create)
	r.Get("/my", details)
	r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Route("/{permission_id}", func(r chi.Router) {
		r.Put("/", update)
		r.Delete("/", delete)
	})

	return r
}

// OrgRequestRouter - Create endpoint for organisation permission request
func OrgRequestRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/", request)

	return r
}
