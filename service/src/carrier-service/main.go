package main

import (
	"carrier-service/api/middleware"
	"carrier-service/api/router"
	"carrier-service/database"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	conn, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	env := NewEnvironment(conn)

	middleware.FiberMiddleware(app)
	handler := env.NewAppHandler()
	router.NewRouter(app, handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "localhost:8080"
	}
	err = app.Listen(port)
	if err != nil {
		log.Fatal(err)
	}

}
