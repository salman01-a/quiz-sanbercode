package router

import (
	"book-management/controllers"
	middlewares "book-management/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/api/users/login", controllers.Login)

	// Protected routes
	protected := r.Group("/", middlewares.BasicAuthMiddleware())

	protected.GET("/api/books", controllers.GetBooks)
	protected.POST("/api/books", controllers.CreateBook)
	//protected.GET("/api/books/:id", controllers.GetBook)
	protected.DELETE("/api/books/:id", controllers.DeleteBook)

	protected.GET("/api/categories", controllers.GetCategories)
	protected.POST("/api/categories", controllers.CreateCategory)
	protected.GET("/api/categories/:id", controllers.GetCategory)
	protected.DELETE("/api/categories/:id", controllers.DeleteCategory)
	protected.GET("/api/categories/:id/books", controllers.GetBooksByCategory)

	return r
}
