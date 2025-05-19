//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	controllers "test-cqrs/src/App/Controllers"
	helpers "test-cqrs/src/App/Helpers"
	commandrepository "test-cqrs/src/App/Repository/CommandRepository"
	commandservice "test-cqrs/src/App/Service/CommandService"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var NewSet = wire.NewSet(
	commandrepository.NewCommandRepositoryImpl,
	commandservice.NewCommandServiceImpl,
	controllers.NewControllerImpl,
	helpers.NewDb,
	NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)),
	NewServer,
)

func InitServer() (*http.Server, func(), error) {
	wire.Build(NewSet)
	return nil, nil, nil
}
