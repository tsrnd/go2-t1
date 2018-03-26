package Helpers

import "net/http"

func InitStaticPrefix() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("app/Views/"))))
}
