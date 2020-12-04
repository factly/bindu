package policy

import (
	"github.com/factly/bindu-server/model"
	"github.com/factly/bindu-server/util"
	"github.com/go-chi/chi"
)

type policyReq struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Permissions []model.Permission `json:"permissions"`
	Users       []string           `json:"users"`
}

// Router - Group of medium router
func Router() chi.Router {
	r := chi.NewRouter()

	entity := "policies"

	r.With(util.CheckKetoPolicy(entity, "get")).Get("/", list)
	r.With(util.CheckKetoPolicy(entity, "create")).Put("/", upsert)
	r.With(util.CheckKetoPolicy(entity, "create")).Post("/default", createDefaults)

	r.Route("/{policy_id}", func(r chi.Router) {
		r.With(util.CheckKetoPolicy(entity, "get")).Get("/", details)
		r.With(util.CheckKetoPolicy(entity, "delete")).Delete("/", delete)
	})

	return r
}
