package orders

import (
	"context"
	"database/sql"
	"time"

	"telegramshop_backend/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repository interface {
	CreateOrder(ctx context.Context, input models.CreateOrder) (models.OrderWithProducts, error)
	GetOrderByID(ctx context.Context, id int) (models.OrderWithProducts, error)
	GetUserOrders(ctx context.Context, userID int64) ([]models.OrderWithProducts, error)
	GetAll(ctx context.Context) ([]models.OrderWithProducts, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateOrder(ctx context.Context, input models.CreateOrder) (models.OrderWithProducts, error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return models.OrderWithProducts{}, err
	}
	defer tx.Rollback()

	orderQuery := `
		INSERT INTO orders (user_id, status)
		VALUES ($1, $2)
		RETURNING id, user_id, status, created_at`

	var order models.OrderWithProducts
	err = tx.QueryRowContext(ctx, orderQuery, input.UserID, "pending").Scan(
		&order.ID, &order.UserID, &order.Status, &order.CreatedAt,
	)
	if err != nil {
		return models.OrderWithProducts{}, err
	}

	productQuery := `
		INSERT INTO order_products (order_id, product_id, quantity, price)
		VALUES ($1, $2, $3, $4)
		RETURNING id, order_id, product_id, quantity, price`

	for _, item := range input.Items {
		var orderProduct models.OrderProduct
		err = tx.QueryRowContext(ctx, productQuery, order.ID, item.ProductID, item.Quantity, 100).Scan(
			&orderProduct.ID, &orderProduct.OrderID, &orderProduct.ProductID, &orderProduct.Quantity, &orderProduct.Price,
		)
		if err != nil {
			return models.OrderWithProducts{}, err
		}
		order.Products = append(order.Products, orderProduct)
	}

	if err = tx.Commit(); err != nil {
		return models.OrderWithProducts{}, err
	}

	return order, nil
}

func (r *repository) GetOrderByID(ctx context.Context, id int) (models.OrderWithProducts, error) {
	orderQuery := `
		SELECT o.id, o.user_id, o.status, o.created_at
		FROM orders o
		WHERE o.id = $1`

	var order models.OrderWithProducts
	err := r.db.GetContext(ctx, &order, orderQuery, id)
	if err == sql.ErrNoRows {
		return models.OrderWithProducts{}, nil
	}
	if err != nil {
		return models.OrderWithProducts{}, err
	}

	productsQuery := `
		SELECT id, order_id, product_id, quantity, price
		FROM order_products
		WHERE order_id = $1`

	var products []models.OrderProduct
	err = r.db.SelectContext(ctx, &products, productsQuery, id)
	if err != nil {
		return models.OrderWithProducts{}, err
	}
	order.Products = products

	return order, nil
}

func (r *repository) GetUserOrders(ctx context.Context, userID int64) ([]models.OrderWithProducts, error) {
	query := `
		SELECT 
			o.id, o.user_id, o.status, o.created_at,
			COALESCE(op.id, 0) as product_id, 
			COALESCE(op.order_id, 0) as product_order_id,
			COALESCE(op.product_id, 0) as product_product_id,
			COALESCE(op.quantity, 0) as product_quantity,
			COALESCE(op.price, 0) as product_price
		FROM orders o
		LEFT JOIN order_products op ON o.id = op.order_id
		WHERE o.user_id = $1
		ORDER BY o.created_at DESC, op.id`

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[int64]*models.OrderWithProducts)
	var orderIDs []int64

	for rows.Next() {
		var (
			orderID, userID                                       int64
			status                                                string
			createdAt                                             time.Time
			productID, productOrderID, productProductID, quantity int
			price                                                 float64
		)

		err := rows.Scan(
			&orderID, &userID, &status, &createdAt,
			&productID, &productOrderID, &productProductID, &quantity, &price,
		)
		if err != nil {
			return nil, err
		}

		if _, exists := ordersMap[orderID]; !exists {
			ordersMap[orderID] = &models.OrderWithProducts{
				ID:        orderID,
				UserID:    userID,
				Status:    status,
				CreatedAt: createdAt,
				Products:  []models.OrderProduct{},
			}
			orderIDs = append(orderIDs, orderID)
		}

		if productID > 0 {
			ordersMap[orderID].Products = append(ordersMap[orderID].Products, models.OrderProduct{
				ID:        productID,
				OrderID:   productOrderID,
				ProductID: productProductID,
				Quantity:  quantity,
				Price:     price,
			})
		}
	}

	result := make([]models.OrderWithProducts, 0, len(orderIDs))
	for _, orderID := range orderIDs {
		result = append(result, *ordersMap[orderID])
	}

	return result, nil
}

func (r *repository) GetAll(ctx context.Context) ([]models.OrderWithProducts, error) {
	query := `
		SELECT 
			o.id, o.user_id, o.status, o.created_at,
			COALESCE(op.id, 0) as product_id, 
			COALESCE(op.order_id, 0) as product_order_id,
			COALESCE(op.product_id, 0) as product_product_id,
			COALESCE(op.quantity, 0) as product_quantity,
			COALESCE(op.price, 0) as product_price
		FROM orders o
		LEFT JOIN order_products op ON o.id = op.order_id
		ORDER BY o.created_at DESC, op.id`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ordersMap := make(map[int64]*models.OrderWithProducts)
	var orderIDs []int64

	for rows.Next() {
		var (
			orderID, userID                                       int64
			status                                                string
			createdAt                                             time.Time
			productID, productOrderID, productProductID, quantity int
			price                                                 float64
		)

		err := rows.Scan(
			&orderID, &userID, &status, &createdAt,
			&productID, &productOrderID, &productProductID, &quantity, &price,
		)
		if err != nil {
			return nil, err
		}

		if _, exists := ordersMap[orderID]; !exists {
			ordersMap[orderID] = &models.OrderWithProducts{
				ID:        orderID,
				UserID:    userID,
				Status:    status,
				CreatedAt: createdAt,
				Products:  []models.OrderProduct{},
			}
			orderIDs = append(orderIDs, orderID)
		}

		if productID > 0 {
			ordersMap[orderID].Products = append(ordersMap[orderID].Products, models.OrderProduct{
				ID:        productID,
				OrderID:   productOrderID,
				ProductID: productProductID,
				Quantity:  quantity,
				Price:     price,
			})
		}
	}

	result := make([]models.OrderWithProducts, 0, len(orderIDs))
	for _, orderID := range orderIDs {
		result = append(result, *ordersMap[orderID])
	}

	return result, nil
}
