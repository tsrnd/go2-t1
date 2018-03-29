package Handlers

import (
	validate "go-t1/app/Http/Validate"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

//
type UserHandler struct {
	repo UserRepository
}

func NewUserHandler(repo UserRepository) *UserHandler {
	return &UserHandler{repo}
}

//
type error interface {
	Error() string
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
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func (controller UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	userId, _ := strconv.ParseUint(r.FormValue("id"), 10, 32)
	controller.repo.DeleteUser(uint32(userId))
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
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

	// validate database
	user := validate.User{
		ErrName:       name,
		ErrCity:       city,
		ErrIdentityID: r.FormValue("identity-id"),
		ErrGender:     gender,
	}
	error := user.Validate()
	if error != nil {
		m := make(map[string]string)
		message := strings.Split(error.Error(), ";")
		for _, value := range message {
			msg := strings.Split(strings.Trim(value, " "), ":")
			m[msg[0]] = value
		}
		validate.Render(w, "app/Views/Users/create.html", m)
	} else {
		controller.repo.InsertUser(name, city, identityID, gender)
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
