package Handlers

import (
	DB "go-t1/Database"
	"go-t1/app/Models"
	"html/template"
	"net/http"
)

type UserHandler struct {
}

//
func (UserHandler) Index(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	tmpl, err := template.ParseFiles("app/Views/Users/index.html")
	if err != nil {
		panic(err)
	}
	rs, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	user := Models.User{}
	users := []Models.User{}
	for rs.Next() {
		err := rs.Scan(&user.Id, &user.Name, &user.City, &user.Indentity_id, &user.Gender)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	tmpl.Execute(w, users)
}

//
func (UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/Views/Users/create.html")
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, nil)
}

//
func (UserHandler) Store(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	db := DB.Connect()
	r.ParseForm()
	name := r.FormValue("name")
	city := r.FormValue("city")
	identityId := r.FormValue("identity-id")
	gender := r.FormValue("gender")
	// Save to database
	stmt, err := db.Prepare(`
		INSERT INTO users(name, city, identity_id, gender)
		VALUES(?, ?, ?, ?)
	`)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	_, err = stmt.Exec(name, city, identityId, gender)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", 301)
}
