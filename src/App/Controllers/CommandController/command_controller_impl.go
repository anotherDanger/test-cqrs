package commandcontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	helpers "test-cqrs/src/App/Helpers"
	commandservice "test-cqrs/src/App/Service/CommandService"
	domain "test-cqrs/src/Domain"
	webapi "test-cqrs/src/WebApi"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type CommandControllerImpl struct {
	svc commandservice.CommandService
}

func NewCommandControllerImpl(svc commandservice.CommandService) CommandController {
	return &CommandControllerImpl{
		svc: svc,
	}
}

func (ctrl *CommandControllerImpl) AddBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	reqBody := &domain.Domain{}

	if reqBody.Author == "" || reqBody.Title == "" || reqBody.Genre == "" {
		helpers.NewErr("/home/andhikadanger/cqrs/src/App/logs/controller", logrus.ErrorLevel, errors.New("body is empty"))
		w.WriteHeader(400)
		return
	}

	json.NewDecoder(r.Body).Decode(reqBody)

	entity, err := ctrl.svc.AddBook(r.Context(), reqBody)
	if err != nil {
		helpers.NewErr("/home/andhikadanger/cqrs/src/App/logs/controller", logrus.ErrorLevel, err)
		return
	}

	response := webapi.Response[*domain.Domain]{
		Code:   201,
		Status: "OK",
		Data:   entity,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response)
}
