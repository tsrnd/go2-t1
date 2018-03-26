package Routes

import (
	. "go-t1/app/Http/Controllers"

	// . "go-t1/app/Helpers"

	"github.com/go-chi/chi"
)

func Routing() *chi.Mux {
	r := chi.NewRouter()

	//define route
	r.Get("/", UserController{}.Index)
	return r
}
