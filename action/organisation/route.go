package organisation

import "github.com/go-chi/chi"

// Router - Group of organisation router
func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/", list)

	return r

}
