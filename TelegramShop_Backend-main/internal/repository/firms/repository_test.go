package firms_test

import (
	"context"
	"telegramshop_backend/internal/repository/firms"
	"testing"

	"telegramshop_backend/internal/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres driver
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=root password=1111 dbname=telegram sslmode=disable")
	require.NoError(t, err)
	return db
}

func TestFirmRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := firms.NewRepository(db)
	ctx := context.Background()

	// Тест создания фирмы
	t.Run("CreateFirm", func(t *testing.T) {
		firm := models.Firm{Name: "Apple"}
		createdFirm, err := repo.CreateFirm(ctx, firm)
		require.NoError(t, err)
		require.NotZero(t, createdFirm.ID)
		require.Equal(t, firm.Name, createdFirm.Name)
	})

	// Тест получения фирмы по ID
	t.Run("GetFirmByID", func(t *testing.T) {
		firm := models.Firm{Name: "Another Firm"}
		createdFirm, err := repo.CreateFirm(ctx, firm)
		require.NoError(t, err)

		gotFirm, err := repo.GetFirmByID(ctx, createdFirm.ID)
		require.NoError(t, err)
		require.Equal(t, createdFirm.ID, gotFirm.ID)
		require.Equal(t, createdFirm.Name, gotFirm.Name)
	})

	// Тест получения всех фирм
	t.Run("GetAllFirms", func(t *testing.T) {
		_, err := repo.CreateFirm(ctx, models.Firm{Name: "Firm1"})
		require.NoError(t, err)
		_, err = repo.CreateFirm(ctx, models.Firm{Name: "Firm2"})
		require.NoError(t, err)

		firms, err := repo.GetAllFirms(ctx)
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(firms), 2)
	})

	// Тест обновления фирмы
	t.Run("UpdateFirm", func(t *testing.T) {
		firm := models.Firm{Name: "Old Name"}
		createdFirm, err := repo.CreateFirm(ctx, firm)
		require.NoError(t, err)

		createdFirm.Name = "New Name"
		err = repo.UpdateFirm(ctx, createdFirm)
		require.NoError(t, err)

		updatedFirm, err := repo.GetFirmByID(ctx, createdFirm.ID)
		require.NoError(t, err)
		require.Equal(t, "New Name", updatedFirm.Name)
	})

	// Тест удаления фирмы
	t.Run("DeleteFirm", func(t *testing.T) {
		firm := models.Firm{Name: "Delete Firm"}
		createdFirm, err := repo.CreateFirm(ctx, firm)
		require.NoError(t, err)

		err = repo.DeleteFirm(ctx, createdFirm.ID)
		require.NoError(t, err)

		deletedFirm, err := repo.GetFirmByID(ctx, createdFirm.ID)
		require.Equal(t, int64(0), deletedFirm.ID) // zero value means not found or empty
	})
}
