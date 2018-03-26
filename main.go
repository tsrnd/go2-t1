package main

import (
	"go-t1/app/Http/Routes"
	"log"
	"net/http"
	"os"

	. "go-t1/app/Helpers"

	"github.com/subosito/gotenv"
)

func main() {
	InitStaticPrefix()
	gotenv.Load()
	router := Routes.Routing()
	log.Fatal(http.ListenAndServe(os.Getenv("HOST_NAME")+":"+os.Getenv("PORT"), router))
}
