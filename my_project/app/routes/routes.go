package routes

import (
	"github.com/gorilla/mux"
	"your-project/app/handler"
)

func SetupRoutes(router *mux.Router, enterpriseHandler *handler.EnterpriseHandler, userHandler *handler.UserHandler) {
	router.HandleFunc("/api/v1/orgs", enterpriseHandler.CreateEnterprise).Methods("POST")
	router.HandleFunc("/api/v1/orgs/{uuid}/users", userHandler.CreateUser).Methods("POST")
}
