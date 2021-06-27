package role

import (
	"github.com/factly/bindu-server/util"
	"github.com/go-chi/chi"
)

type roleReq struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description"`
	Users       []string `json:"users"`
}

// Router - Group of medium router
func Router() chi.Router {
	r := chi.NewRouter()

	entity := "roles"

	r.With(util.CheckKetoPolicy(entity, "get")).Get("/", list)
	r.With(util.CheckKetoPolicy(entity, "update")).Put("/", upsert)

	r.Route("/{role_id}", func(r chi.Router) {
		r.With(util.CheckKetoPolicy(entity, "get")).Get("/", details)
		r.With(util.CheckKetoPolicy(entity, "delete")).Delete("/", delete)
	})

	return r
}
