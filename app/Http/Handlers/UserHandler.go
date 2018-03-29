package Handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

func (controller UserHandler) Edit(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.ParseUint(chi.URLParam(r, "user"), 10, 32)
	user := controller.repo.Show(uint32(userID))
	if user.ID == 0 {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	tmpl, _ := template.ParseFiles("app/Views/Users/details.html")
	tmpl.Execute(w, user)
}

func (controller UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	userID, _ := strconv.ParseUint(chi.URLParam(r, "user"), 10, 32)
	name := r.FormValue("name")
	city := r.FormValue("city")
	identityID, _ := strconv.ParseInt(r.FormValue("identityID"), 10, 64)
	gender, _ := strconv.ParseBool(r.FormValue("gender"))
	userData := map[string]interface{}{
		"name":        name,
		"city":        city,
		"indentityId": identityID,
		"gender":      gender,
	}
	controller.repo.Update(uint32(userID), userData)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
