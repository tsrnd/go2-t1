package main

import (
	DB "go-t1/Database"
	"go-t1/app/Http/Routes"
	"log"
	"net/http"
	"os"

	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	db := DB.Connect()
	router := Routes.Routing(db)
	log.Fatal(http.ListenAndServe(os.Getenv("HOST_NAME")+":"+os.Getenv("PORT"), router))
}
