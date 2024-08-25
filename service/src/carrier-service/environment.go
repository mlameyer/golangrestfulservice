package main

import (
	"carrier-service/api/handler"
	"carrier-service/domain/repository"
	"carrier-service/infrastructure"
	"carrier-service/usecase"
	"gorm.io/gorm"
)

type Environment interface {
	NewCarrierRepository() repository.CarrierRepository
	NewCarrierUseCase() usecase.CarrierUseCase
	NewCarrierHandler() handler.CarrierHandler
	NewLoginHandler() handler.LoginHandler
	NewAppHandler() handler.AppHandler
}

type environment struct {
	Conn *gorm.DB
}

func NewEnvironment(Conn *gorm.DB) Environment {
	return &environment{Conn}
}

type appHandler struct {
	handler.CarrierHandler
	handler.LoginHandler
}

func (e *environment) NewAppHandler() handler.AppHandler {
	appHandler := &appHandler{}
	appHandler.CarrierHandler = e.NewCarrierHandler()
	appHandler.LoginHandler = e.NewLoginHandler()
	return appHandler
}

func (e *environment) NewCarrierRepository() repository.CarrierRepository {
	return infrastructure.NewCarrierRepository(e.Conn)
}

func (e *environment) NewCarrierUseCase() usecase.CarrierUseCase {
	return usecase.NewCarrierUseCase(e.NewCarrierRepository())
}

func (e *environment) NewCarrierHandler() handler.CarrierHandler {
	return handler.NewCarrierHandler(e.NewCarrierUseCase())
}

func (e *environment) NewLoginHandler() handler.LoginHandler {
	return handler.NewLoginHandler()
}
