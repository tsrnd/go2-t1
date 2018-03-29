package Handlers

import (
	// . "go-t1/app/Models"

	"html/template"
	"net/http"
	"strconv"
)

type UserHandler struct {
	repo UserRepository
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{repo}
}

//
func (controller UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/Views/Users/index.html")
	users := controller.repo.GetListUser()
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	tmpl.Execute(w, users)
}

func (controller UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseUint(r.FormValue("id"), 10, 32)
	controller.repo.DeleteUser(uint32(userId))
	http.Redirect(w, r, "/", 301)
}
