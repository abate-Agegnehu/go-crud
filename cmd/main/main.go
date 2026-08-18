package main

import (
	"fmt"
	"go-crud/internal/database"
	"go-crud/internal/handler"
	"go-crud/internal/repository"
	"go-crud/internal/service"
	"log"
	"net/http"
)

func main() {
	// Connect to database
	err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer database.DB.Close()

	// Initialize repository, service, and handler
	userRepo := repository.NewUserRepository(database.DB.Conn)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Define routes
	http.HandleFunc("/api/users/create", userHandler.Create)
	http.HandleFunc("/api/users/all", userHandler.GetAll)
	http.HandleFunc("/api/users/get", userHandler.GetByID)
	http.HandleFunc("/api/users/update", userHandler.Update)
	http.HandleFunc("/api/users/delete", userHandler.Delete)

	// Start server
	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
