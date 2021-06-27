package space

import (
	"net/http"

	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/middlewarex"
	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type spacePermissionRequest struct {
	Title       string         `json:"title" validate:"required"`
	Description postgres.Jsonb `json:"description" swaggertype:"primitive,string"`
	Charts      int64          `json:"charts"`
	SpaceID     int64          `json:"space_id" validate:"required"`
}

var permissionContext config.ContextKey = "space_perm_user"
var requestUserContext config.ContextKey = "request_user"

// Router - CRUD servies
func Router() http.Handler {
	r := chi.NewRouter()

	app := "bindu"

	r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Get("/", list)
	r.Get("/my", my)
	r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Route("/{request_id}", func(r chi.Router) {
		r.Get("/", details)
		r.Delete("/", delete)
		r.Post("/approve", approve)
		r.Post("/reject", reject)
	})

	return r
}
