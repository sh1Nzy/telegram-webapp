package handler

import (
	"telegramshop_backend/internal/service/avg_marks"
	"telegramshop_backend/internal/service/basket"
	"telegramshop_backend/internal/service/categories"
	"telegramshop_backend/internal/service/comment"
	"telegramshop_backend/internal/service/favorites"
	"telegramshop_backend/internal/service/firms"
	"telegramshop_backend/internal/service/marks"
	"telegramshop_backend/internal/service/orders"
	"telegramshop_backend/internal/service/prices"
	"telegramshop_backend/internal/service/products"
	"telegramshop_backend/internal/service/users"

	"github.com/gofiber/fiber/v2"
)

var (
	apiPath     = "/api"
	versionPath = "v1"
	basePath    = apiPath + "/" + versionPath
)

type Handler struct {
	userService     users.Service
	favoriteService favorites.Service
	basketService   basket.Service
	orderService    orders.Service
	firmsService    firms.Service
	priceService    prices.Service
	categoryService categories.Service
	productService  products.Service
	marksService    marks.MarksService
	avgMarksService avg_marks.AvgMarksService
	commentService  comment.CommentService
}

func NewHandler(
	userService users.Service,
	favoriteService favorites.Service,
	basketService basket.Service,
	orderService orders.Service,
	firmsService firms.Service,
	priceService prices.Service,
	categoryService categories.Service,
	productService products.Service,
	marksService marks.MarksService,
	avgMarksService avg_marks.AvgMarksService,
	commentService comment.CommentService,
) *Handler {
	return &Handler{
		userService:     userService,
		favoriteService: favoriteService,
		basketService:   basketService,
		orderService:    orderService,
		firmsService:    firmsService,
		priceService:    priceService,
		categoryService: categoryService,
		productService:  productService,
		marksService:    marksService,
		avgMarksService: avgMarksService,
		commentService:  commentService,
	}
}

func (h *Handler) InitRouter(app *fiber.App) {
	api := app.Group(basePath)

	// User routes
	api.Post("/users", h.CreateUser)
	api.Get("/users", h.GetAllUsers)
	api.Get("/users/:id", h.GetUser)
	api.Delete("/users/:id", h.DeleteUser)

	// Favorites routes
	api.Post("/favorites", h.AddToFavorites)
	api.Get("/favorites/:user_id", h.GetUserFavorites)
	api.Delete("/favorites/:user_id/:product_id", h.RemoveFromFavorites)

	// Basket routes
	api.Post("/basket", h.AddToBasket)
	api.Get("/basket/:user_id", h.GetUserBasket)
	api.Put("/basket", h.UpdateBasketItem)
	api.Delete("/basket/:user_id/:product_id", h.RemoveFromBasket)

	// Orders routes
	api.Post("/orders", h.CreateOrder)
	api.Get("/orders/:id", h.GetOrder)
	api.Get("/orders/user/:user_id", h.GetUserOrders)
	api.Get("/orders/all", h.GetAllOrders)

	//firms
	api.Post("/firms", h.CreateFirm)       //work
	api.Get("/firms/:id", h.GetFirmByID)   //work
	api.Get("/firms", h.GetAllFirms)       //work
	api.Put("/firms/:id", h.UpdateFirm)    //work
	api.Delete("/firms/:id", h.DeleteFirm) //work

	//price
	api.Post("/prices", h.CreatePrice)                                   //work
	api.Get("/prices/:id", h.GetPriceByID)                               //work
	api.Get("/prices/product/:product_id", h.GetPricesByProductID)       //work
	api.Put("/prices/:id", h.UpdatePrice)                                //work
	api.Delete("/prices/:id", h.DeletePrice)                             //work
	api.Delete("/prices/product/:product_id", h.DeletePricesByProductID) //work
	api.Patch("/prices/:id/count", h.UpdatePriceCount)

	// category
	api.Post("/categories", h.CreateCategory)                  //work
	api.Get("/categories/:id", h.GetCategoryByID)              //work
	api.Get("/categories", h.GetAllCategories)                 //work
	api.Put("/categories/:id", h.UpdateCategory)               //work
	api.Delete("/categories/:id", h.DeleteCategory)            //work
	api.Put("/categories/:id/image", h.SetCategoryImage)       //work
	api.Delete("/categories/:id/image", h.RemoveCategoryImage) //work

	// product
	api.Post("/products", h.CreateProduct)       //work
	api.Get("/products/:id", h.GetProductByID)   //work
	api.Get("/products", h.GetAllProducts)       //work
	api.Put("/products/:id", h.UpdateProduct)    //work
	api.Delete("/products/:id", h.DeleteProduct) //work

	// product images
	api.Put("/products/:id/image", h.AddProductImage)       //work
	api.Delete("/products/:id/image", h.RemoveProductImage) //work в теле запроса нужно указать адрес удаляемой img
	api.Put("/products/:id/images", h.SetProductImages)     //work

	// product stats
	api.Patch("/products/:id/sell", h.IncrementSellCount) //work
	api.Patch("/products/:id/stock", h.UpdateStock)       //work

	api.Get("/marks/user/:user_id", h.GetUserMarks)                           ///work
	api.Get("/marks/user/:user_id/product/:product_id", h.GetProductUserMark) //work
	api.Post("/marks/user/:user_id/product/:product_id", h.AddMark)           //work
	api.Delete("/marks/user/:user_id/product/:product_id", h.DeleteMark)      //work

	api.Get("/avg_marks/product/:product_id", h.GetAvgMark) //work
	api.Get("/avg_marks", h.GetAllAvgMarks)                 //work

	api.Post("/comments/user/:user_id/product/:product_id", h.AddComment)      //work
	api.Put("/comments/user/:user_id/product/:product_id", h.EditComment)      //work
	api.Delete("/comments/user/:user_id/product/:product_id", h.DeleteComment) //work
	api.Get("/comments/product/:product_id", h.GetCommentsByProduct)           //work
}
