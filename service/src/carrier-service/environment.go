package main

import (
	"carrier-service/api/handler"
	"carrier-service/usecase"
	"database/sql"
)

type Environment interface {
	NewCarrierRepository() repository.CarrierRepository
	NewCarrierUseCase() usecase.CarrierUseCase
	NewCarrierHandler() handler.CarrierHandler
	NewAppHandler() handler.AppHandler
}

type environment struct {
	Conn *sql.DB
}

func NewEnvironment(Conn *sql.DB) Environment {
	return &environment{Conn}
}

type appHandler struct {
	handler.CarrierHandler
}

func (e *environment) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.CarrierHandler = e.NewCarrierHandler()
	return appHandler
}

func (e *environment) NewCarrierRepository() repository.CarrierRepository {
	return datastore.NewUserRepository(e.Conn)
}

func (e *environment) NewCarrierUseCase() usecase.CarrierUseCase {
	return usecase.NewCarrierUseCase(e.NewUserRepository())
}

func (e *environment) NewCarrierHandler() handler.CarrierHandler {
	return handler.NewCarrierHandler(e.NewCarrierUseCase())
}
