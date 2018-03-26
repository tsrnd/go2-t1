package Helpers

import (
	"net/http"
)

func AuthenticateMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if GetSession("authenticated", r) != true {
			http.Redirect(w, r, "/login", http.StatusPermanentRedirect)
		}
		next(w, r)
	}
}
