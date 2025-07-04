package handler

import (
	"strconv"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

// CreateOrder creates a new order
// @Summary Create new order
// @Description Creates a new order for a user with specified products
// @Tags orders
// @Accept json
// @Produce json
// @Param order body models.CreateOrder true "Order creation data"
// @Success 200 {object} models.OrderResponse "Order successfully created"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/orders [post]
func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	var input models.CreateOrder
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	order, err := h.orderService.CreateOrder(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_create_order", err.Error()))
	}

	return c.JSON(web.OkResp("success_order_created", order))
}

// GetOrder retrieves order by ID
// @Summary Get order by ID
// @Description Returns order details with all products
// @Tags orders
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.OrderResponse "Order retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid order ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/orders/{id} [get]
func (h *Handler) GetOrder(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_order_id", "Invalid order ID"))
	}

	order, err := h.orderService.GetOrderByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_order", err.Error()))
	}

	return c.JSON(web.OkResp("success_order_retrieved", order))
}

// GetUserOrders retrieves all orders for a specific user
// @Summary Get user's orders
// @Description Returns all orders for a specific user
// @Tags orders
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} models.OrderListResponse "User's orders retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid user ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/orders/user/{user_id} [get]
func (h *Handler) GetUserOrders(c *fiber.Ctx) error {
	userIDStr := c.Params("user_id")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_user_id", "Invalid user ID"))
	}

	orders, err := h.orderService.GetUserOrders(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_user_orders", err.Error()))
	}

	return c.JSON(web.OkResp("success_user_orders_retrieved", orders))
}

// GetAllOrders retrieves all orders
// @Summary Get all orders
// @Description Returns all orders in the system
// @Tags orders
// @Produce json
// @Success 200 {object} models.OrderListResponse "All orders retrieved successfully"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/orders/all [get]
func (h *Handler) GetAllOrders(c *fiber.Ctx) error {
	orders, err := h.orderService.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_all_orders", err.Error()))
	}

	return c.JSON(web.OkResp("success_all_orders_retrieved", orders))
}
