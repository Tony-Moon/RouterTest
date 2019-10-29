package main

import (
	"net/http"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("page2").HTTPBox()))
	http.ListenAndServe(":9000", router)

}
