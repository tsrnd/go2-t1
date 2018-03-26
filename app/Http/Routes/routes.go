package Routes

import (
	. "go-t1/app/Http/Controllers"

	. "go-t1/app/Helpers"

	"github.com/go-chi/chi"
)

func Routing() *chi.Mux {
	r := chi.NewRouter()

	//define route
	r.Get("/", AuthenticateMiddleware(UserController{}.Index))
	r.Get("/register", UserController{}.Register)
	r.Post("/register", UserController{}.Create)
	r.Get("/login", UserController{}.LoginForm)
	r.Post("/login", UserController{}.Login)
	return r
}
