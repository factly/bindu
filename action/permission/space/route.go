package space

import (
	"github.com/factly/bindu-server/config"
	"github.com/factly/bindu-server/util"
	"github.com/go-chi/chi"
)

type spacePermission struct {
	SpaceID uint  `json:"space_id" validate:"required"`
	Charts  int64 `json:"charts"`
}

var userContext config.ContextKey = "space_perm_user"

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
