package main

import (
	"database/sql"
	"log"
	"net/http"

	"nexttalenta-backend/database"
	"nexttalenta-backend/services"
	loginhttp "nexttalenta-backend/transport/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Initialize database connection
	dsn := "root@tcp(localhost:3306)/nexttalenta"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Check database connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Database is unreachable: %v", err)
	}

	log.Println("Database connected successfully.")

	// ✅ Initialize repository and use case correctly
	userRepo := database.NewUserRepository(db)
	loginService := services.NewLoginService(userRepo)
	loginHandler := loginhttp.NewLoginHandler(loginService) // ✅ Updated

	// Set up HTTP routes
	http.HandleFunc("/login", loginHandler.Login)

	// Start server
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
