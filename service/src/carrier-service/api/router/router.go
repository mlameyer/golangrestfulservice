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
	api.Get("/:id", middleware.JWTProtected(), h.GetCarrier)
}
