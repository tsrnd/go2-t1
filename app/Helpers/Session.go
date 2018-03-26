package Helpers

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func SetSession(key string, value interface{}, w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	session.Values[key] = value
	session.Save(r, w)
}

func GetSession(key string, r *http.Request) interface{} {
	session, _ := store.Get(r, "cookie-name")
	return session.Values[key]
}
