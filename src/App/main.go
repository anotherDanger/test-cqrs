package main

import (
	"fmt"
	"net/http"
	controllers "test-cqrs/src/App/Controllers"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(ctrl controllers.Controller) *httprouter.Router {
	r := httprouter.New()
	r.POST("/v1/books", ctrl.AddBook)
	return r
}

func NewServer(handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
}

func main() {
	server, cleanup, err := InitServer()
	if err != nil {
		fmt.Println(err)
	}

	defer cleanup()

	server.ListenAndServe()
}
