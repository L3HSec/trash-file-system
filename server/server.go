package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func RegisterAPI(method string, path string, handle http.HandlerFunc) {
	router.Handle(path, handle).Methods(method)
	fmt.Printf("API %s method %s registered\n", path, method)
}

//Run launch the http server
func Run() {
	//TODO fix relative path
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("../static"))))
	http.ListenAndServe(":8080", router)
}

func init() {
	router = mux.NewRouter()
}
