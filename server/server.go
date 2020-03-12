package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

//RegisterAPI register api urls
func RegisterAPI(method string, path string, handle http.HandlerFunc) {
	router.Handle(path, handle).Methods(method)
	fmt.Printf("API %s method %s registered\n", path, method)
}

//Run launch the http server
func Run(addr string) {
	fmt.Println("Server listening at " + addr)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServe(addr, router)
	if err != nil {
		fmt.Println(err)
	}
}

func init() {
	router = mux.NewRouter()
}
