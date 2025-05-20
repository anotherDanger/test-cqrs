package commandcontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CommandController interface {
	AddBook(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
