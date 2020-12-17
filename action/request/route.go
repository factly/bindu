package request

import (
	"net/http"

	"github.com/factly/bindu-server/action/request/organisationPermission"
	"github.com/factly/bindu-server/action/request/spacePermission"
	"github.com/go-chi/chi"
)

// Router - CRUD servies
func Router() http.Handler {
	r := chi.NewRouter()

	r.Mount("/space-permissions", spacePermission.Router())
	r.Mount("/organisation-permissions", organisationPermission.Router())

	return r
}
