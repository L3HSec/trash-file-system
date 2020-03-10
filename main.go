package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	fmt.Println("server launched")
	router := httprouter.New()
	router.ServeFiles("/*filepath", http.Dir("static"))
	http.ListenAndServe(":8080", router)
}
