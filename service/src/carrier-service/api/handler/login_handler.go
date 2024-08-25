package handler

import (
	"carrier-service/api/dto"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"time"
)

type LoginHandler interface {
	Authenticate(c *fiber.Ctx) error
}

type loginHandler struct {
	jwtToken string
}

func NewLoginHandler() LoginHandler {
	return &loginHandler{}
}

func (l loginHandler) Authenticate(c *fiber.Ctx) error {
	request := new(dto.AuthenticationRequest)

	if err := c.BodyParser(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := request.User
	pass := request.Password

	// Throws Unauthorized error
	if user != "jack" || pass != "burton" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"name": "Jack Burton",
		"exp":  time.Now().Add(time.Minute * 120).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"token": t})
}
