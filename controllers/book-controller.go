package controllers

import (
	"book-management/utils"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	rows, err := utils.DB.Query("SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching books"})
		return
	}
	defer rows.Close()

	books := []map[string]interface{}{}
	for rows.Next() {
		var book map[string]interface{}
		var id, releaseYear, price, totalPage, categoryID int
		var title, description, imageURL, thickness, createdBy, modifiedBy sql.NullString
		var createdAt, modifiedAt sql.NullTime

		err := rows.Scan(&id, &title, &description, &imageURL, &releaseYear, &price, &totalPage, &thickness, &categoryID, &createdAt, &createdBy, &modifiedAt, &modifiedBy)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error scanning book"})
			return
		}

		book = map[string]interface{}{
			"id":           id,
			"title":        title.String,
			"description":  description.String,
			"image_url":    imageURL.String,
			"release_year": releaseYear,
			"price":        price,
			"total_page":   totalPage,
			"thickness":    thickness.String,
			"category_id":  categoryID,
			"created_at":   createdAt.Time,
			"created_by":   createdBy.String,
			"modified_at":  modifiedAt.Time,
			"modified_by":  modifiedBy.String,
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func CreateBook(c *gin.Context) {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		ImageURL    string `json:"image_url"`
		ReleaseYear int    `json:"release_year"`
		Price       int    `json:"price"`
		TotalPage   int    `json:"total_page"`
		CategoryID  int    `json:"category_id"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	thickness := "tipis"
	if input.TotalPage > 100 {
		thickness = "tebal"
	}

	query := `INSERT INTO books (title, description, image_url, release_year, price, total_page, thickness, category_id, created_by) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 'admin') RETURNING id`
	var bookID int
	err := utils.DB.QueryRow(query, input.Title, input.Description, input.ImageURL, input.ReleaseYear, input.Price, input.TotalPage, thickness, input.CategoryID).Scan(&bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": bookID})
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	var book map[string]interface{}
	var title, description, imageURL, thickness, createdBy, modifiedBy sql.NullString
	var releaseYear, price, totalPage, categoryID int
	var createdAt, modifiedAt sql.NullTime

	query := "SELECT id, title, description, image_url, release_year, price, total_page, thickness, category_id, created_at, created_by, modified_at, modified_by FROM books WHERE id=$1"
	err = utils.DB.QueryRow(query, id).Scan(&id, &title, &description, &imageURL, &releaseYear, &price, &totalPage, &thickness, &categoryID, &createdAt, &createdBy, &modifiedAt, &modifiedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching book"})
		return
	}

	book = map[string]interface{}{
		"id":           id,
		"title":        title.String,
		"description":  description.String,
		"image_url":    imageURL.String,
		"release_year": releaseYear,
		"price":        price,
		"total_page":   totalPage,
		"thickness":    thickness.String,
		"category_id":  categoryID,
		"created_at":   createdAt.Time,
		"created_by":   createdBy.String,
		"modified_at":  modifiedAt.Time,
		"modified_by":  modifiedBy.String,
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid book ID"})
		return
	}

	query := "DELETE FROM books WHERE id=$1"
	res, err := utils.DB.Exec(query, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting book"})
		return
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching rows affected"})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}
