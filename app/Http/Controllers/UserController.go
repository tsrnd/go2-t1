package Controllers

import (
	DB "go-t1/Database"
	"go-t1/app/Models"
	"html/template"
	"net/http"
)

type UserController struct {
}

//
func (UserController) Index(w http.ResponseWriter, r *http.Request) {
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
