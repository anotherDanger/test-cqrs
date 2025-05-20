//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	commandcontroller "test-cqrs/src/App/Controllers/CommandController"
	querycontroller "test-cqrs/src/App/Controllers/QueryController"
	helpers "test-cqrs/src/App/Helpers"
	commandrepository "test-cqrs/src/App/Repository/CommandRepository"
	queryrepository "test-cqrs/src/App/Repository/QueryRepository"
	commandservice "test-cqrs/src/App/Service/CommandService"
	queryservice "test-cqrs/src/App/Service/QueryService"

	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var NewSet = wire.NewSet(
	commandrepository.NewCommandRepositoryImpl,
	commandservice.NewCommandServiceImpl,
	commandcontroller.NewCommandControllerImpl,
	queryrepository.NewQueryRepositoryImpl,
	queryservice.NewQueryServiceImpl,
	querycontroller.NewQueryControllerImpl,
	helpers.NewDb,
	NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)),
	NewServer,
)

func InitServer() (*http.Server, func(), error) {
	wire.Build(NewSet)
	return nil, nil, nil
}
