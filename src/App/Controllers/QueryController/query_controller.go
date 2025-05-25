package querycontroller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type QueryController interface {
	GetBook(w http.ResponseWriter, r *http.Request, p httprouter.Params)
	GetBookByTitle(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}
