package main

import (
	"log"
	"net/http"

	"backend-app/config"
	"backend-app/internal/handler"
	"backend-app/internal/repository"
	"backend-app/internal/usecase"
	"backend-app/pkg/database"
	"backend-app/routes"
)

func main() {
	config.LoadEnv()


	db := database.ConnectMySQL()
	defer db.Close()

	r := routes.NewRouter(userHandler)

	
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
