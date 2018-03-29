package Validate

import (
	"html/template"
	"net/http"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

//
type User struct {
	Name       string
	City       string
	IdentityID string
	Gender     bool
}

func (a User) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
		validation.Field(&a.City, validation.Required),
		validation.Field(&a.IdentityID, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{9}$")).Error("Must be a string with night digits")),
		validation.Field(&a.Gender, validation.Required),
	)
}

//
func Render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
