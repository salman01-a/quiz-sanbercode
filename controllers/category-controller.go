package controllers

import (
	"net/http"

	"book-management/models"
	"book-management/services"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	categories, err := services.GetCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func CreateCategory(c *gin.Context) {
	var input models.Category
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input"})
		return
	}

	categoryID, err := services.CreateCategory(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error creating category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": categoryID, "message": "Category created"})
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	category, err := services.GetCategoryByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching category"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func DeleteCategory(c *gin.Context) {
	id := c.Param("id")
	err := services.DeleteCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error deleting category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}

func GetBooksByCategory(c *gin.Context) {
	id := c.Param("id")
	books, err := services.GetBooksByCategory(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error fetching books"})
		return
	}

	c.JSON(http.StatusOK, books)
}
