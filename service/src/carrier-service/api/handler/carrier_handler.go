package handler

import (
	"carrier-service/usecase"
	"github.com/gofiber/fiber/v2"
	"strconv"
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
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	if carrier == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": true,
			"message": "Carrier not found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Successfully fetched carrier",
		"carrier": carrier,
	})
}
