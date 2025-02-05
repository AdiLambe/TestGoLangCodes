package main

import (
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/handler"
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/repository"
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/routes"
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize in-memory repositories
	enterpriseRepo := repository.NewEnterpriseRepository()
	userRepo := repository.NewUserRepository()

	// Initialize services
	enterpriseService := service.NewEnterpriseService(enterpriseRepo)
	userService := service.NewUserService(userRepo)

	// Initialize handlers
	enterpriseHandler := handler.NewEnterpriseHandler(enterpriseService)
	userHandler := handler.NewUserHandler(userService)

	// Setup routes
	router := mux.NewRouter()
	routes.SetupRoutes(router, enterpriseHandler, userHandler)

	// Start server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
