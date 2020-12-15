package permission

import (
	"github.com/factly/bindu-server/action/permission/organisation"
	"github.com/factly/bindu-server/action/permission/space"
	"github.com/go-chi/chi"
)

// Router router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Mount("/organisations", organisation.Router())
	r.Mount("/spaces", space.Router())

	return r

}
