package querycontroller

import (
	"encoding/json"
	"net/http"
	helpers "test-cqrs/src/App/Helpers"
	queryservice "test-cqrs/src/App/Service/QueryService"
	domain "test-cqrs/src/Domain"
	webapi "test-cqrs/src/WebApi"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type QueryControllerImpl struct {
	svc queryservice.QueryService
}

func NewQueryControllerImpl(svc queryservice.QueryService) QueryController {
	return &QueryControllerImpl{
		svc: svc,
	}
}

func (ctrl *QueryControllerImpl) GetBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	key := p.ByName("key")
	value := p.ByName("value")

	result, err := ctrl.svc.GetBook(r.Context(), key, value)
	if err != nil {
		helpers.NewErr("/app/src/App/logs/controller", logrus.ErrorLevel, err)
		return
	}

	response := webapi.Response[[]*domain.Domain]{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)
}
