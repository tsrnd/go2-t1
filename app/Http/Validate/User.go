package Validate

import (
	"html/template"
	"net/http"
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
)

//
type User struct {
	ErrName       string
	ErrCity       string
	ErrIdentityID string
	ErrGender     bool
}

func (a User) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.ErrName, validation.Required),
		validation.Field(&a.ErrCity, validation.Required),
		validation.Field(&a.ErrIdentityID, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{9}$")).Error("Must be a string with nine digits")),
		validation.Field(&a.ErrGender, validation.Required),
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
