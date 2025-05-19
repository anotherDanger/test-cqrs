package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Controller interface {
	AddBook(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
