package handler

import (
	"encoding/json"
	"net/http"
	"your-project/app/dto"
	"your-project/app/service"
)

type UserHandler struct {
	Service *service.UserService
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	json.NewDecoder(r.Body).Decode(&req)

	user, err := h.Service.CreateUser(req.UserName, req.Email, req.EnterpriseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
