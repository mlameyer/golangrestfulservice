package router

import (
	"carrier-service/api/handler"
	"carrier-service/api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewRouter(app *fiber.App, h handler.AppHandler) {
	api := app.Group("/api", logger.New())
	api.Post("/authenticate", h.Authenticate)

	api.Get("/carriers/:id", middleware.JWTProtected(), h.GetCarrier)
	api.Post("/carriers", middleware.JWTProtected(), h.CreateCarrier)
	api.Put("/carriers/:id/address", middleware.JWTProtected(), h.UpdateCarrierAddress)
	api.Put("/carriers/:id/active", middleware.JWTProtected(), h.UpdateCarrierActiveStatus)
	api.Delete("/carriers/:id", middleware.JWTProtected(), h.DeleteCarrier)
}
