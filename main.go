package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	router "book-management/routers"
	"book-management/utils"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %v", err)
	}

	// Initialize the database connection
	err := initDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Run migrations
	err = runMigrations()
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	// Set up the Gin router
	r := router.SetupRouter()

	// Define the port from environment variables or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	log.Printf("Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// initDB initializes the database connection
func initDB() error {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	utils.DB = db
	return nil
}

// runMigrations applies database migrations
func runMigrations() error {
	migrations := &migrate.FileMigrationSource{
		Dir: "migrations",
	}

	_, err := migrate.Exec(utils.DB, "postgres", migrations, migrate.Up)
	return err
}
