// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"test-cqrs/src/App/Controllers/CommandController"
	"test-cqrs/src/App/Controllers/QueryController"
	"test-cqrs/src/App/Helpers"
	"test-cqrs/src/App/Repository/CommandRepository"
	"test-cqrs/src/App/Repository/QueryRepository"
	"test-cqrs/src/App/Service/CommandService"
	"test-cqrs/src/App/Service/QueryService"
)

// Injectors from injector.go:

func InitServer() (*http.Server, func(), error) {
	db, cleanup, err := helpers.NewDb()
	if err != nil {
		return nil, nil, err
	}
	commandRepository := commandrepository.NewCommandRepositoryImpl()
	commandService := commandservice.NewCommandServiceImpl(db, commandRepository)
	commandController := commandcontroller.NewCommandControllerImpl(commandService)
	queryRepository := queryrepository.NewQueryRepositoryImpl()
	queryService := queryservice.NewQueryServiceImpl(queryRepository)
	queryController := querycontroller.NewQueryControllerImpl(queryService)
	router := NewRouter(commandController, queryController)
	server := NewServer(router)
	return server, func() {
		cleanup()
	}, nil
}

// injector.go:

var NewSet = wire.NewSet(commandrepository.NewCommandRepositoryImpl, commandservice.NewCommandServiceImpl, commandcontroller.NewCommandControllerImpl, queryrepository.NewQueryRepositoryImpl, queryservice.NewQueryServiceImpl, querycontroller.NewQueryControllerImpl, helpers.NewDb, NewRouter, wire.Bind(new(http.Handler), new(*httprouter.Router)), NewServer)
