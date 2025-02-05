package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/domain"
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/service"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	Service *service.UserService
}

// NewUserHandler initializes the handler
func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// CreateUser handles user creation
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Extract enterprise ID from the URL
	vars := mux.Vars(r)
	enterpriseID, err := uuid.Parse(vars["enterpriseID"])
	if err != nil {
		http.Error(w, "Invalid enterprise ID", http.StatusBadRequest)
		return
	}
	user.EnterpriseID = enterpriseID

	// Create the user
	if err := h.Service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// GetAllUsers retrieves all users for a given enterprise
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	enterpriseID, err := uuid.Parse(vars["enterpriseID"])
	if err != nil {
		http.Error(w, "Invalid enterprise ID", http.StatusBadRequest)
		return
	}

	users, err := h.Service.GetAllUsers(enterpriseID)
	if err != nil {
		http.Error(w, "Failed to retrieve users", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
