package handler

import (
	"carrier-service/api/dto"
	"carrier-service/api/middleware"
	"carrier-service/domain/model"
	"carrier-service/usecase"
	"github.com/gofiber/fiber/v2"
	"strconv"
	"sync"
	"time"
)

type CarrierHandler interface {
	GetCarrier(c *fiber.Ctx) error
	CreateCarrier(c *fiber.Ctx) error
	UpdateCarrierAddress(c *fiber.Ctx) error
	UpdateCarrierActiveStatus(c *fiber.Ctx) error
	DeleteCarrier(c *fiber.Ctx) error
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

func (h *carrierHandler) CreateCarrier(c *fiber.Ctx) error {
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

	carrierDto := &dto.CreateCarrier{}
	if err := c.BodyParser(carrierDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	carrier := &model.Carrier{}
	err = carrier.NewCarrier(
		carrierDto.Name,
		carrierDto.Address,
		carrierDto.Active,
	)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	ctx := c.Context()

	e := make(chan error)
	var wg sync.WaitGroup
	wg.Add(1)
	go h.CarrierUseCase.CreateCarrier(ctx, carrier, e, &wg)
	wg.Wait()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
		"message": "Successfully created carrier",
	})
}

func (h *carrierHandler) UpdateCarrierAddress(c *fiber.Ctx) error {
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

	addressDto := &dto.UpdateCarrierAddress{}
	if err := c.BodyParser(addressDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ctx := c.Context()
	carrier, err := h.CarrierUseCase.GetCarrier(ctx, addressDto.Id)
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

	err = carrier.UpdateCarrierAddress(addressDto.Address)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	err = h.CarrierUseCase.UpdateCarrier(ctx, carrier)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Successfully updated address",
	})
}

func (h *carrierHandler) UpdateCarrierActiveStatus(c *fiber.Ctx) error {
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

	statusDto := &dto.UpdateCarrierActivityStatus{}
	if err := c.BodyParser(statusDto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	ctx := c.Context()
	carrier, err := h.CarrierUseCase.GetCarrier(ctx, statusDto.Id)
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

	carrier.UpdateCarrierActiveStatus(statusDto.Active)

	err = h.CarrierUseCase.UpdateCarrier(ctx, carrier)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Successfully updated carrier active status",
	})
}

func (h *carrierHandler) DeleteCarrier(c *fiber.Ctx) error {
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

	carrier := &model.Carrier{}
	carrier.ID = id

	ctx := c.Context()
	err = h.CarrierUseCase.DeleteCarrier(ctx, carrier)
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
		"message": "Successfully deleted carrier",
	})
}
