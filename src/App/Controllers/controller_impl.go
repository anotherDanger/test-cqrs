package controllers

import (
	"encoding/json"
	"net/http"
	helpers "test-cqrs/src/App/Helpers"
	commandservice "test-cqrs/src/App/Service/CommandService"
	domain "test-cqrs/src/Domain"
	webapi "test-cqrs/src/WebApi"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type ControllerImpl struct {
	svc commandservice.CommandService
}

func NewControllerImpl(svc commandservice.CommandService) Controller {
	return &ControllerImpl{}
}

func (ctrl *ControllerImpl) AddBook(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	reqBody := &domain.Domain{}

	json.NewDecoder(r.Body).Decode(reqBody)

	entity, err := ctrl.svc.AddBook(r.Context(), reqBody)
	if err != nil {
		helpers.NewErr("../logs/controller", logrus.ErrorLevel, err)
		return
	}

	response := webapi.Response[domain.Domain]{
		Code:   201,
		Status: "OK",
		Data:   *entity,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(response)
}
