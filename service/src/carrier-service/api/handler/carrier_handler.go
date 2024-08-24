package handler

import (
	"carrier-service/api/middleware"
	"carrier-service/usecase"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"time"
)

type CarrierHandler interface {
	GetCarrier(c *fiber.Ctx) error
}

type carrierHandler struct {
	CarrierUseCase usecase.CarrierUseCase
}

func NewCarrierHandler(c usecase.CarrierUseCase) CarrierHandler {
	return &carrierHandler{c}
}

func (h *carrierHandler) GetCarrier(c *fiber.Ctx) error {
	now := time.Now().Unix()

	// Get claims from JWT.
	claims, err := middleware.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	expires := claims.Expires

	if now > expires {
		// Return status 401 and unauthorized error message.
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": true,
			"msg":   "unauthorized, check expiration time of your token",
		})
	}

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	ctx := c.Context()
	carrier, err := h.CarrierUseCase.GetCarrier(ctx, id)
	if err != nil {
		if err.Error() == "record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"success": true,
				"message": "Carrier not found",
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Successfully fetched carrier",
		"carrier": carrier,
	})
}
