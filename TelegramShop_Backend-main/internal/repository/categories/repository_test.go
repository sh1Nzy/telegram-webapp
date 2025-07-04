package categories_test

import (
	"context"
	"testing"

	"telegramshop_backend/internal/models"
	"telegramshop_backend/internal/repository/categories"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Connect("postgres", "host=localhost port=5432 user=root password=1111 dbname=telegram sslmode=disable")
	require.NoError(t, err)
	return db
}

func TestCategoryRepository(t *testing.T) {
	db := setupTestDB(t)
	repo := categories.NewRepository(db)
	ctx := context.Background()

	// Тест создания категории
	t.Run("CreateCategory", func(t *testing.T) {
		category := models.Category{Name: "Clothing"}
		createdCategory, err := repo.CreateCategory(ctx, category)
		require.NoError(t, err)
		require.NotZero(t, createdCategory.ID)
		require.Equal(t, category.Name, createdCategory.Name)
	})

	t.Run("GetCategoryByID", func(t *testing.T) {
		category := models.Category{Name: "Accessories"}
		createdCategory, err := repo.CreateCategory(ctx, category)
		require.NoError(t, err)

		gotCategory, err := repo.GetCategoryByID(ctx, createdCategory.ID)
		require.NoError(t, err)
		require.Equal(t, createdCategory.ID, gotCategory.ID)
		require.Equal(t, createdCategory.Name, gotCategory.Name)
	})

	t.Run("GetAllCategories", func(t *testing.T) {
		_, err := repo.CreateCategory(ctx, models.Category{Name: "Books"})
		require.NoError(t, err)
		_, err = repo.CreateCategory(ctx, models.Category{Name: "Toys"})
		require.NoError(t, err)

		categories, err := repo.GetAllCategories(ctx)
		require.NoError(t, err)
		require.GreaterOrEqual(t, len(categories), 2)

	})

	t.Run("UpdateCategory", func(t *testing.T) {
		category := models.Category{Name: "Old Category"}
		createdCategory, err := repo.CreateCategory(ctx, category)
		require.NoError(t, err)

		createdCategory.Name = "Updated Category"
		err = repo.UpdateCategory(ctx, createdCategory)
		require.NoError(t, err)

		updatedCategory, err := repo.GetCategoryByID(ctx, createdCategory.ID)
		require.NoError(t, err)
		require.Equal(t, "Updated Category", updatedCategory.Name)
	})

	t.Run("SetAndRemoveImage", func(t *testing.T) {
		category := models.Category{Name: "With Image1"}
		createdCategory, err := repo.CreateCategory(ctx, category)
		require.NoError(t, err)

		imageURL := "https://example.com/image.jpg"
		err = repo.SetImage(ctx, createdCategory.ID, imageURL)
		require.NoError(t, err)

		categoryWithImage, err := repo.GetCategoryByID(ctx, createdCategory.ID)
		require.NoError(t, err)
		require.NotNil(t, categoryWithImage.Image)
		require.Equal(t, imageURL, *categoryWithImage.Image)

		err = repo.RemoveImage(ctx, createdCategory.ID)
		require.NoError(t, err)

		categoryWithoutImage, err := repo.GetCategoryByID(ctx, createdCategory.ID)
		require.NoError(t, err)
		require.Nil(t, categoryWithoutImage.Image)
	})

	t.Run("DeleteCategory", func(t *testing.T) {
		category := models.Category{Name: "To Be Deleted"}
		createdCategory, err := repo.CreateCategory(ctx, category)
		require.NoError(t, err)

		err = repo.DeleteCategory(ctx, createdCategory.ID)
		require.NoError(t, err)

		deletedCategory, err := repo.GetCategoryByID(ctx, createdCategory.ID)
		require.NoError(t, err)
		require.Equal(t, int64(0), deletedCategory.ID)
	})
}
