package main

import (
	"carrier-service/database"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	conn, err := database.NewConnection()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(fmt.Sprintf("Failed to close: %v", err))
		}
	}()
	env := NewEnvironment(conn)
	handler := env.NewAppHandler()

	router.NewRouter(env, handler)
	middleware.NewMiddleware(e)
	if err := e.Start(fmt.Sprintf(":%d", conf.Current.Server.Port)); err != nil {
		e.Logger.Fatal(fmt.Sprintf("Failed to start: %v", err))
	}

	app := fiber.New()

	app.Use(middleware.Logger())

	router.SetupRoutes(app)

	app.Listen(3000)
}
