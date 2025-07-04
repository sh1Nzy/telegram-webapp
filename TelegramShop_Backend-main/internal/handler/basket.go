package handler

import (
	"strconv"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

// AddToBasket adds item to user's basket
// @Summary Add item to basket
// @Description Adds a product to user's basket
// @Tags basket
// @Accept json
// @Produce json
// @Param item body models.BasketItem true "Basket item data"
// @Success 200 {object} models.BasketResponse "Item successfully added to basket"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/basket [post]
func (h *Handler) AddToBasket(c *fiber.Ctx) error {
	var input models.BasketItem
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	item, err := h.basketService.AddToBasket(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_add_to_basket", err.Error()))
	}

	return c.JSON(web.OkResp("success_item_added_to_basket", item))
}

// GetUserBasket retrieves user's basket
// @Summary Get user's basket
// @Description Returns all items in user's basket
// @Tags basket
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} models.BasketListResponse "User's basket retrieved"
// @Failure 400 {object} models.ErrorResponse "Invalid user ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/basket/{user_id} [get]
func (h *Handler) GetUserBasket(c *fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_user_id", "Invalid user ID"))
	}

	items, err := h.basketService.GetUserBasket(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_user_basket", err.Error()))
	}

	return c.JSON(web.OkResp("success_user_basket_retrieved", items))
}

// UpdateBasketItem updates item in user's basket
// @Summary Update basket item
// @Description Updates quantity of product in user's basket
// @Tags basket
// @Accept json
// @Produce json
// @Param item body models.BasketItem true "Updated basket item data"
// @Success 200 {object} models.BasketResponse "Basket item successfully updated"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/basket [put]
func (h *Handler) UpdateBasketItem(c *fiber.Ctx) error {
	var input models.BasketItem
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	item, err := h.basketService.UpdateBasketItem(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_update_basket_item", err.Error()))
	}

	return c.JSON(web.OkResp("success_basket_item_updated", item))
}

// RemoveFromBasket removes item from user's basket
// @Summary Remove item from basket
// @Description Removes product from user's basket
// @Tags basket
// @Produce json
// @Param user_id path int true "User ID"
// @Param product_id path int true "Product ID"
// @Success 200 {object} models.SuccessResponse "Item successfully removed from basket"
// @Failure 400 {object} models.ErrorResponse "Invalid parameters"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/basket/{user_id}/{product_id} [delete]
func (h *Handler) RemoveFromBasket(c *fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	productIDStr := c.Params("product_id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_user_id", "Invalid user ID"))
	}

	productID, err := strconv.Atoi(productIDStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_product_id", "Invalid product ID"))
	}

	if err := h.basketService.RemoveFromBasket(c.Context(), userID, productID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_remove_from_basket", err.Error()))
	}

	return c.JSON(web.OkResp("success_item_removed_from_basket", nil))
}
