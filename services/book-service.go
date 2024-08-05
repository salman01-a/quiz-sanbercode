package services

import (
	"database/sql"
	"errors"

	"book-management/models"
	"book-management/utils"
)

func GetBooks() ([]models.Book, error) {
	rows, err := utils.DB.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []models.Book{}
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func CreateBook(input models.Book) (int, error) {
	var bookID int

	thickness := "tipis"
	if input.TotalPage > 100 {
		thickness = "tebal"
	}

	query := `
        INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), 'admin', NOW(), 'admin') RETURNING id`

	err := utils.DB.QueryRow(query, input.Title, input.Description, input.ImageURL, input.ReleaseYear, input.Price, input.TotalPage, thickness, input.CategoryID).Scan(&bookID)
	if err != nil {
		return 0, err
	}

	return bookID, nil
}

func GetBookByID(id string) (models.Book, error) {
	var book models.Book
	query := "SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE id=$1"
	err := utils.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Description, &book.ImageURL, &book.ReleaseYear, &book.Price, &book.TotalPage, &book.Thickness, &book.CategoryID, &book.CreatedAt, &book.CreatedBy, &book.ModifiedAt, &book.ModifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return book, errors.New("book not found")
		}
		return book, err
	}

	return book, nil
}

func DeleteBook(id string) error {
	query := "DELETE FROM books WHERE id=$1"
	result, err := utils.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no book found to delete")
	}

	return nil
}
