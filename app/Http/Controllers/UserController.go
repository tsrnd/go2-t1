package Controllers

import (
	"fmt"
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
	tmpl, err := template.ParseFiles("app/Views/sample.html")
	if err != nil {
		panic(err)
	}
	rs, err := db.Query("SELECT * FROM employee")
	if err != nil {
		panic(err.Error())
	}
	user := Models.User{}
	users := []Models.User{}
	for rs.Next() {
		err := rs.Scan(&user.Id, &user.Name, &user.City)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	tmpl.Execute(w, users)
}

//
func (UserController) Register(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/Views/Users/register.html")
	if err != nil {
		panic(err.Error())
	}
	tmpl.Execute(w, nil)
}

func (UserController) LoginForm(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("app/Views/Users/login.html")
	if err != nil {
		panic(err.Error())
	}
	tmpl.Execute(w, nil)
}

func (UserController) Login(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
	}
	insForm, err := db.Prepare("SELECT * employee WHERE username=? AND password=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(r.FormValue("username"), r.FormValue("password"))
	fmt.Println(r.FormValue("username"))
}

//
func (UserController) Create(w http.ResponseWriter, r *http.Request) {
	db := DB.Connect()
	defer db.Close()
	if r.Method != "POST" {
		http.Error(w, "Not found", http.StatusNotFound)
	}
	insForm, err := db.Prepare("INSERT INTO employee(username, password) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(r.FormValue("username"), r.FormValue("password"))
	http.Redirect(w, r, "/login", http.StatusMovedPermanently)
}
