package Routes

import (
	. "go-t1/app/Http/Handlers"

	// . "go-t1/app/Helpers"

	"github.com/go-chi/chi"
)

func Routing() *chi.Mux {
	r := chi.NewRouter()

	//define route
	r.Get("/", UserHandler{}.Index)
	r.Get("/add", UserHandler{}.Create)
	r.Post("/add", UserHandler{}.Store)
	return r
}
