package space

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type spacePermission struct {
	SpaceID uint  `json:"space_id" validate:"required"`
	Charts  int64 `json:"charts"`
}

type spacePermissionRequest struct {
	Title       string         `json:"title" validate:"required"`
	Description postgres.Jsonb `json:"description" swaggertype:"primitive,string"`
	Charts      int64          `json:"charts"`
	SpaceID     int64          `json:"space_id" validate:"required"`
}

var userContext config.ContextKey = "space_perm_user"
var requestUserContext config.ContextKey = "request_user"

// Router - Group of currency router
func Router() chi.Router {
	r := chi.NewRouter()

	r.With(util.CheckSuperOrganisation).Get("/", list)
	r.Get("/my", my)
	r.With(util.CheckSuperOrganisation).Post("/", create)
	r.With(util.CheckSuperOrganisation).Route("/{permission_id}", func(r chi.Router) {
		r.Get("/", details)
		r.Put("/", update)
		r.Delete("/", delete)
	})

	return r
}

// SpaceRequestRouter - Create endpoint for space permission request
func SpaceRequestRouter() chi.Router {
	r := chi.NewRouter()
	r.Post("/", request)

	return r
}
