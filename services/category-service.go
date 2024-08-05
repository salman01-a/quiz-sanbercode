package services

import (
	"book-management/models"
	"book-management/utils"
	"database/sql"
	"errors"
)

func GetCategories() ([]models.Category, error) {
	rows, err := utils.DB.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func CreateCategory(category models.Category) (int, error) {
	var categoryID int
	query := "INSERT INTO categories (name, created_by, modified_by) VALUES ($1, $2, $3) RETURNING id"
	err := utils.DB.QueryRow(query, category.Name, category.CreatedBy, category.ModifiedBy).Scan(&categoryID)
	if err != nil {
		return 0, err
	}

	return categoryID, nil
}

func GetCategoryByID(id string) (models.Category, error) {
	var category models.Category
	query := "SELECT * FROM categories WHERE id=$1"
	err := utils.DB.QueryRow(query, id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.CreatedBy, &category.ModifiedAt, &category.ModifiedBy)
	if err == sql.ErrNoRows {
		return category, errors.New("category not found")
	}
	if err != nil {
		return category, err
	}

	return category, nil
}

func DeleteCategory(id string) error {
	query := "DELETE FROM categories WHERE id=$1"
	result, err := utils.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("category not found")
	}

	return nil
}

func GetBooksByCategory(categoryID string) ([]models.Book, error) {
	rows, err := utils.DB.Query("SELECT * FROM books WHERE category_id=$1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}
