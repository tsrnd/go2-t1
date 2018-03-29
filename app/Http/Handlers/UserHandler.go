package Handlers

import (
	validate "go-t1/app/Http/Validate"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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
		Name:       name,
		City:       city,
		IdentityID: r.FormValue("identity-id"),
		Gender:     gender,
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
		http.Redirect(w, r, "/", 301)
	}
}
