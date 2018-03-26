package Helpers

import (
	"html/template"
	"net/http"
)

func View(path string, w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("app/Views/Users/register.html")
	if err != nil {
		panic(err.Error())
	}
	tmpl.Execute(w, nil)
}
