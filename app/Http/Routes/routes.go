package Routes

import (
	. "go-t1/app/Http/Handlers"

	"github.com/go-chi/chi"
	"github.com/jinzhu/gorm"
)

func Routing(db *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	//define route
	r.Get("/", userHandler(db).Index)

	return r
}

func userHandler(db *gorm.DB) *UserHandler {
	repo := NewUserRepository(db)
	return NewUserHandler(repo)
}
