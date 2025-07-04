package handler

import (
	"strconv"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/pkg/web"

	"github.com/gofiber/fiber/v2"
)

// CreateProduct creates a new product
// @Summary Create new product
// @Description Creates a new product with specified details
// @Tags products
// @Accept json
// @Produce json
// @Param product body models.Product true "Product creation data"
// @Success 200 {object} models.ProductResponse "Product successfully created"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products [post]
func (h *Handler) CreateProduct(c *fiber.Ctx) error {
	var input models.Product
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	product, err := h.productService.CreateProduct(c.Context(), input)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_create_product", err.Error()))
	}

	return c.JSON(web.OkResp("success_product_created", product))
}

// GetProductByID retrieves product by ID
// @Summary Get product by ID
// @Description Returns product details by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.Product "Product retrieved successfully"
// @Failure 400 {object} models.ErrorResponse "Invalid product ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id} [get]
func (h *Handler) GetProductByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid product ID"))
	}

	product, err := h.productService.GetProductByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_product", err.Error()))
	}

	return c.JSON(web.OkResp("success_product_retrieved", product))
}

// GetAllProducts retrieves all products
// @Summary Get all products
// @Description Returns all products in the system
// @Tags products
// @Produce json
// @Success 200 {object} models.ProductListResponse "All products retrieved successfully"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products [get]
func (h *Handler) GetAllProducts(c *fiber.Ctx) error {
	products, err := h.productService.GetAllProducts(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_get_all_products", err.Error()))
	}

	return c.JSON(web.OkResp("success_products_retrieved", products))
}

// UpdateProduct updates product details
// @Summary Update product
// @Description Updates product details by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body models.UpdateProductInput true "Updated product data"
// @Success 200 {object} models.SuccessResponse "Product successfully updated"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id} [put]
func (h *Handler) UpdateProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid product ID"))
	}
	var input models.UpdateProductInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_request_body", "Invalid request body"))
	}

	if err := h.productService.UpdateProduct(c.Context(), id, input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_update_product", err.Error()))
	}

	return c.JSON(web.OkResp("success_product_updated", nil))
}

// DeleteProduct deletes a product
// @Summary Delete product
// @Description Deletes a product by its ID
// @Tags products
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.SuccessResponse "Product successfully deleted"
// @Failure 400 {object} models.ErrorResponse "Invalid product ID"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id} [delete]
func (h *Handler) DeleteProduct(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid product ID"))
	}

	if err := h.productService.DeleteProduct(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_delete_product", err.Error()))
	}

	return c.JSON(web.OkResp("success_product_deleted", nil))
}

// AddProductImage adds image to product
// @Summary Add product image
// @Description Adds a new image to a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param image body models.ImagesInput true "Image data"
// @Success 200 {object} models.SuccessResponse "Product image successfully added"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id}/images [post]
func (h *Handler) AddProductImage(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid product ID"))
	}

	var input models.ImageInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_body", "Invalid body"))
	}

	if err := h.productService.AddProductImage(c.Context(), id, input.Image); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_add_product_image", err.Error()))
	}

	return c.JSON(web.OkResp("success_product_image_added", nil))
}

// RemoveProductImage removes image from product
// @Summary Remove product image
// @Description Removes an image from a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param image body models.ImagesInput true "Image data to remove"
// @Success 200 {object} models.SuccessResponse "Product image successfully removed"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id}/images [delete]
func (h *Handler) RemoveProductImage(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid product ID"))
	}

	var input models.ImageInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_body", "Invalid body"))
	}

	if err := h.productService.RemoveProductImage(c.Context(), id, input.Image); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_remove_product_image", err.Error()))
	}

	return c.JSON(web.OkResp("success_product_image_removed", nil))
}

// SetProductImages sets all product images
// @Summary Set product images
// @Description Sets all images for a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param images body models.ImagesInput true "Array of image data"
// @Success 200 {object} models.SuccessResponse "Product images successfully set"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id}/images [put]
func (h *Handler) SetProductImages(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid product ID"))
	}

	var input models.ImagesInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_body", "Invalid body"))
	}

	if err := h.productService.SetProductImages(c.Context(), id, input.Images); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_set_product_images", err.Error()))
	}

	return c.JSON(web.OkResp("success_product_images_set", nil))
}

// IncrementSellCount increments product sell count
// @Summary Increment sell count
// @Description Increments the sell count for a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param count body models.CountInput true "Count to increment"
// @Success 200 {object} models.SuccessResponse "Sell count successfully incremented"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id}/sell-count [put]
func (h *Handler) IncrementSellCount(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid product ID"))
	}

	var input models.CountInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_body", "Invalid body"))
	}

	if err := h.productService.IncrementSellCount(c.Context(), id, input.Count); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_increment_sell_count", err.Error()))
	}

	return c.JSON(web.OkResp("success_sell_count_incremented", nil))
}

// UpdateStock updates product stock
// @Summary Update stock
// @Description Updates the stock count for a product
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param stock body models.StockInput true "New stock value"
// @Success 200 {object} models.SuccessResponse "Stock successfully updated"
// @Failure 400 {object} models.ErrorResponse "Invalid request body"
// @Failure 500 {object} models.ErrorResponse "Internal server error"
// @Router /api/v1/products/{id}/stock [put]
func (h *Handler) UpdateStock(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_id", "Invalid product ID"))
	}

	var input models.StockInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(web.ErrorResp("error_invalid_body", "Invalid body"))
	}

	if err := h.productService.UpdateStock(c.Context(), id, input.Stock); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(web.ErrorResp("error_update_stock", err.Error()))
	}

	return c.JSON(web.OkResp("success_stock_updated", nil))
}
