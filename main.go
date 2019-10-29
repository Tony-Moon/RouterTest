package main

import (
	"net/http"
	"text/template"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("website").HTTPBox()))

	http.HandleFunc("/2", page2)
	http.ListenAndServe(":8080", router)

}

func page2(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	chat, _ := template.ParseFiles("page2/page2.html")
	chat.Execute(res, req)
}
