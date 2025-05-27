package main

import (
	"fmt"
	"net/http"
	commandcontroller "test-cqrs/src/App/Controllers/CommandController"
	querycontroller "test-cqrs/src/App/Controllers/QueryController"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(c_ctrl commandcontroller.CommandController, q_ctrl querycontroller.QueryController) *httprouter.Router {
	r := httprouter.New()
	r.POST("/v1/books", c_ctrl.AddBook)
	r.POST("/v1/books/search/:key/:value", q_ctrl.GetBook)

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
