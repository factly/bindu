package space

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/factly/x/middlewarex"
	"github.com/go-chi/chi"
)

type spacePermission struct {
	SpaceID uint  `json:"space_id" validate:"required"`
	Charts  int64 `json:"charts"`
}

var userContext config.ContextKey = "space_perm_user"
var requestUserContext config.ContextKey = "request_user"

// Router - Group of currency router
func Router() chi.Router {
	r := chi.NewRouter()

	app := "bindu"

	r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Get("/", list)
	r.Get("/my", my)
	r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Post("/", create)
	r.With(middlewarex.CheckSuperOrganisation(app, util.GetOrganisation)).Route("/{permission_id}", func(r chi.Router) {
		r.Get("/", details)
		r.Put("/", update)
		r.Delete("/", delete)
	})

	return r
}
