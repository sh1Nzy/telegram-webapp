package prices_test

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/prices"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=root password=1111 dbname=telegram sslmode=disable")
	require.NoError(t, err)
	return db
}

func TestPriceRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := prices.NewRepository(db)
	ctx := context.Background()

	t.Run("CreateAndGetPriceByID", func(t *testing.T) {
		price := models.Price{
			ProductID: 22,
			Count:     10,
			Price:     999.99,
		}
		created, err := repo.CreatePrice(ctx, price)
		require.NoError(t, err)
		require.NotZero(t, created.ID)

		got, err := repo.GetPriceByID(ctx, created.ID)
		require.NoError(t, err)
		require.Equal(t, created.ID, got.ID)
		require.Equal(t, price.ProductID, got.ProductID)
		require.Equal(t, price.Count, got.Count)
		require.Equal(t, price.Price, got.Price)
	})

	t.Run("UpdatePrice", func(t *testing.T) {
		price := models.Price{
			ProductID: 23,
			Count:     5,
			Price:     499.99,
		}
		created, err := repo.CreatePrice(ctx, price)
		require.NoError(t, err)

		input := models.UpdatePriceInput{
			Price: 999.99,
		}
		err = repo.UpdatePrice(ctx, created.ID, input)
		require.NoError(t, err)

		updated, err := repo.GetPriceByID(ctx, created.ID)
		require.NoError(t, err)
		require.Equal(t, input.Price, updated.Price)
	})

	t.Run("GetPricesByProductID", func(t *testing.T) {
		productID := int64(24)
		for i := 0; i < 3; i++ {
			_, err := repo.CreatePrice(ctx, models.Price{
				ProductID: productID,
				Count:     1 + i,
				Price:     float64(100 + i),
			})
			require.NoError(t, err)
		}

		pricesList, err := repo.GetPricesByProductID(ctx, productID)
		require.NoError(t, err)
		require.Len(t, pricesList, 3)
		for _, p := range pricesList {
			require.Equal(t, productID, p.ProductID)
		}
	})

	t.Run("DeletePrice", func(t *testing.T) {
		price := models.Price{ProductID: 22, Count: 2, Price: 200}
		created, err := repo.CreatePrice(ctx, price)
		require.NoError(t, err)

		err = repo.DeletePrice(ctx, created.ID)
		require.NoError(t, err)

		got, err := repo.GetPriceByID(ctx, created.ID)
		require.NoError(t, err)
		require.Zero(t, got.ID)
	})

	t.Run("DeletePricesByProductID", func(t *testing.T) {
		productID := int64(34)
		for i := 0; i < 2; i++ {
			_, err := repo.CreatePrice(ctx, models.Price{
				ProductID: productID,
				Count:     i + 1,
				Price:     float64(150 + i),
			})
			require.NoError(t, err)
		}

		err := repo.DeletePricesByProductID(ctx, productID)
		require.NoError(t, err)

		pricesList, err := repo.GetPricesByProductID(ctx, productID)
		require.NoError(t, err)
		require.Len(t, pricesList, 0)
	})
	t.Run("UpdatePriceCount", func(t *testing.T) {
		price := models.Price{
			ProductID: 30,
			Count:     10,
			Price:     1000,
		}
		created, err := repo.CreatePrice(ctx, price)
		require.NoError(t, err)

		newCount := 20
		err = repo.UpdatePriceCount(ctx, created.ID, newCount)
		require.NoError(t, err)

		updated, err := repo.GetPriceByID(ctx, created.ID)
		require.NoError(t, err)
		require.Equal(t, newCount, updated.Count)
		require.Equal(t, price.Price, updated.Price)
	})
}
