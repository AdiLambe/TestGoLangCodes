package main

import (
	"net/http"
	"your-project/config"

	"github.com/gorilla/mux"
)

func main() {
	config.ConnectDatabase()
	router := mux.NewRouter()
	http.ListenAndServe(":8080", router)
}
