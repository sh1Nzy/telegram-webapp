package handler

import (
	"strconv"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

// CreatePrice creates a new price
// @Summary Create new price
// @Description Creates a new price with specified details
// @Tags prices
// @Accept json
// @Produce json
// @Param price body models.Price true "Price creation data"
// @Success 200 {object} models.PriceResponse "Price successfully created"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/prices [post]
func (h *Handler) CreatePrice(c *fiber.Ctx) error {
	var input models.Price
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}
	if input.Count < 0 || input.Price < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_data", "Count and Price must be non-negative"))
	}
	price, err := h.priceService.CreatePrice(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_create_price", err.Error()))
	}

	return c.JSON(web.OkResp("success_price_created", price))
}

// GetPriceByID retrieves price by ID
// @Summary Get price by ID
// @Description Returns price details by its ID
// @Tags prices
// @Produce json
// @Param id path int true "Price ID"
// @Success 200 {object} models.PriceResponse "Price retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid price ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/prices/{id} [get]
func (h *Handler) GetPriceByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid price ID"))
	}

	price, err := h.priceService.GetPriceByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_price", err.Error()))
	}

	return c.JSON(web.OkResp("success_price_retrieved", price))
}

// GetPricesByProductID retrieves prices by product ID
// @Summary Get prices by product ID
// @Description Returns all prices for a specific product
// @Tags prices
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {object} models.PriceListResponse "Prices retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid product ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/prices/product/{product_id} [get]
func (h *Handler) GetPricesByProductID(c *fiber.Ctx) error {
	productIDStr := c.Params("product_id")
	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_product_id", "Invalid product ID"))
	}

	prices, err := h.priceService.GetPricesByProductID(c.Context(), productID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_prices_by_product", err.Error()))
	}

	return c.JSON(web.OkResp("success_prices_retrieved", prices))
}

// UpdatePrice updates price details
// @Summary Update price
// @Description Updates price details by its ID
// @Tags prices
// @Accept json
// @Produce json
// @Param id path int true "Price ID"
// @Param price body models.UpdatePriceInput true "Updated price data"
// @Success 200 {object} models.SuccessResponse "Price successfully updated"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/prices/{id} [put]
func (h *Handler) UpdatePrice(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid price ID"))
	}
	var input models.UpdatePriceInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	if err := h.priceService.UpdatePrice(c.Context(), id, input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_update_price", err.Error()))
	}

	return c.JSON(web.OkResp("success_price_updated", nil))
}

// DeletePrice deletes a price
// @Summary Delete price
// @Description Deletes a price by its ID
// @Tags prices
// @Produce json
// @Param id path int true "Price ID"
// @Success 200 {object} models.SuccessResponse "Price successfully deleted"
// @Failure 400 {object} models.ErrorResponse "Invalid price ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/prices/{id} [delete]
func (h *Handler) DeletePrice(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid price ID"))
	}

	if err := h.priceService.DeletePrice(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_delete_price", err.Error()))
	}

	return c.JSON(web.OkResp("success_price_deleted", nil))
}

// DeletePricesByProductID deletes all prices for a product
// @Summary Delete prices by product ID
// @Description Deletes all prices associated with a specific product
// @Tags prices
// @Produce json
// @Param product_id path int true "Product ID"
// @Success 200 {object} models.SuccessResponse "Prices successfully deleted"
// @Failure 400 {object} models.ErrorResponse "Invalid product ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/prices/product/{product_id} [delete]
func (h *Handler) DeletePricesByProductID(c *fiber.Ctx) error {
	productIDStr := c.Params("product_id")
	productID, err := strconv.ParseInt(productIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_product_id", "Invalid product ID"))
	}

	if err := h.priceService.DeletePricesByProductID(c.Context(), productID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_delete_prices_by_product", err.Error()))
	}

	return c.JSON(web.OkResp("success_prices_deleted", nil))
}

// UpdatePriceCount updates price count
// @Summary Update price count
// @Description Updates the count for a specific price
// @Tags prices
// @Accept json
// @Produce json
// @Param id path int true "Price ID"
// @Param count body models.UpdatePriceCount true "New count value"
// @Success 200 {object} models.SuccessResponse "Price count successfully updated"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/prices/{id}/count [put]
func (h *Handler) UpdatePriceCount(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid price ID"))
	}

	var input models.UpdatePriceCount
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	if err := h.priceService.UpdatePriceCount(c.Context(), id, input.NewCount); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_update_price_count", err.Error()))
	}

	return c.JSON(web.OkResp("success_price_count_updated", nil))
}
