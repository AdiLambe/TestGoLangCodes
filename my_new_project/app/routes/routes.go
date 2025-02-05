package routes

import (
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/handler"

	"github.com/gorilla/mux"
)

// SetupRoutes configures API routes
func SetupRoutes(router *mux.Router, enterpriseHandler *handler.EnterpriseHandler, userHandler *handler.UserHandler) {
	// Enterprise routes
	router.HandleFunc("/api/v1/orgs", enterpriseHandler.CreateEnterprise).Methods("POST")
	router.HandleFunc("/api/v1/orgs", enterpriseHandler.GetAllEnterprises).Methods("GET")

	// User routes
	router.HandleFunc("/api/v1/orgs/{enterpriseID}/users", userHandler.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/orgs/{enterpriseID}/users", userHandler.GetAllUsers).Methods("GET")
}
