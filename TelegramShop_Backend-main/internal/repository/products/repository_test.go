package products_test

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/products"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=root password=1111 dbname=telegram sslmode=disable")
	require.NoError(t, err)
	return db
}

func TestProductRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := products.NewRepository(db)
	ctx := context.Background()

	t.Run("CreateProduct", func(t *testing.T) {
		product := models.Product{
			Name:        "Test Product2",
			FirmID:      25,
			Description: "Test Desc",
			CategoryID:  46,
			Attributes:  map[string]interface{}{"size": "M", "test": "rest"},
			SellCount:   5,
			Stock:       10,
			Image:       []string{"https://example.com/image1.jpg", "https://example.com/image2.jpg"},
		}
		created, err := repo.CreateProduct(ctx, product)
		require.NoError(t, err)
		require.NotZero(t, created.ID)
		require.Equal(t, product.Name, created.Name)
	})

	t.Run("GetProductByID", func(t *testing.T) {
		product := models.Product{
			Name:        "Read Product",
			FirmID:      26,
			Description: "Some Desc",
			CategoryID:  47,
		}
		created, err := repo.CreateProduct(ctx, product)
		require.NoError(t, err)

		got, err := repo.GetProductByID(ctx, created.ID)
		require.NoError(t, err)
		require.Equal(t, created.ID, got.ID)
		require.Equal(t, product.Name, got.Name)
	})

	t.Run("UpdateProduct", func(t *testing.T) {
		product := models.Product{
			Name:        "Old Name",
			FirmID:      26,
			Description: "Old Desc",
			CategoryID:  47,
		}
		created, err := repo.CreateProduct(ctx, product)
		require.NoError(t, err)

		created.Name = "Updated Name"
		created.Description = "New Desc"
		err = repo.UpdateProduct(ctx, created)
		require.NoError(t, err)

		updated, err := repo.GetProductByID(ctx, created.ID)
		require.NoError(t, err)
		require.Equal(t, "Updated Name", updated.Name)
		require.Equal(t, "New Desc", updated.Description)
	})

	t.Run("AddAndRemoveImage", func(t *testing.T) {
		product := models.Product{Name: "Image Product", FirmID: 26, CategoryID: 47}
		created, err := repo.CreateProduct(ctx, product)
		require.NoError(t, err)

		image := "https://example.com/image.jpg"
		err = repo.AddProductImage(ctx, created.ID, image)
		require.NoError(t, err)

		updated, err := repo.GetProductByID(ctx, created.ID)
		require.NoError(t, err)
		require.Contains(t, updated.Image, image)

		err = repo.RemoveProductImage(ctx, created.ID, image)
		require.NoError(t, err)

		afterRemove, err := repo.GetProductByID(ctx, created.ID)
		require.NoError(t, err)
		for _, img := range afterRemove.Image {
			require.NotEqual(t, image, img)
		}
	})

	t.Run("SetProductImages", func(t *testing.T) {
		product := models.Product{Name: "Set Image Product", FirmID: 26, CategoryID: 47}
		created, err := repo.CreateProduct(ctx, product)
		require.NoError(t, err)

		images := []string{"img1.jpg", "img2.jpg"}
		err = repo.SetProductImages(ctx, created.ID, images)
		require.NoError(t, err)

		got, err := repo.GetProductByID(ctx, created.ID)
		require.NoError(t, err)
		require.ElementsMatch(t, images, got.Image)
	})

	t.Run("DeleteProduct", func(t *testing.T) {
		product := models.Product{Name: "To Be Deleted", FirmID: 26, CategoryID: 47}
		created, err := repo.CreateProduct(ctx, product)
		require.NoError(t, err)

		err = repo.DeleteProduct(ctx, created.ID)
		require.NoError(t, err)

		got, err := repo.GetProductByID(ctx, created.ID)
		require.NoError(t, err)
		require.Equal(t, int64(0), got.ID)
	})

	t.Run("IncrementSellCount", func(t *testing.T) {
		product := models.Product{
			Name:        "Sell Counter Product",
			FirmID:      26,
			CategoryID:  47,
			SellCount:   0,
			Description: "Test sell count",
		}
		created, err := repo.CreateProduct(ctx, product)
		require.NoError(t, err)

		err = repo.IncrementSellCount(ctx, created.ID, 3)
		require.NoError(t, err)

		got, err := repo.GetProductByID(ctx, created.ID)
		require.NoError(t, err)
		require.Equal(t, 3, got.SellCount)
	})

	t.Run("UpdateStock", func(t *testing.T) {
		product := models.Product{
			Name:        "Stock Update Product",
			FirmID:      26,
			CategoryID:  47,
			Stock:       5,
			Description: "Test stock update",
		}
		created, err := repo.CreateProduct(ctx, product)
		require.NoError(t, err)

		err = repo.UpdateStock(ctx, created.ID, 42)
		require.NoError(t, err)

		got, err := repo.GetProductByID(ctx, created.ID)
		require.NoError(t, err)
		require.Equal(t, 42, got.Stock)
	})
}
