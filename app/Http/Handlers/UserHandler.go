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

//
func (controller UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/Views/Users/create.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, nil)
}

//
func (controller UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	r.ParseForm()
	name := r.FormValue("name")
	city := r.FormValue("city")
	identityID, _ := strconv.ParseInt(r.FormValue("identity-id"), 10, 64)
	gender, _ := strconv.ParseBool(r.FormValue("gender"))
	controller.repo.InsertUser(name, city, identityID, gender)

	http.Redirect(w, r, "/", 301)
}
