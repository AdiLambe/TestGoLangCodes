package handler

import (
	"encoding/json"
	"net/http"

	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/domain"
	"github.com/AdiLambe/TestGoLangCodes/my_new_project/app/service"
)

// EnterpriseHandler handles enterprise-related HTTP requests
type EnterpriseHandler struct {
	Service *service.EnterpriseService
}

// NewEnterpriseHandler initializes the handler
func NewEnterpriseHandler(service *service.EnterpriseService) *EnterpriseHandler {
	return &EnterpriseHandler{Service: service}
}

// CreateEnterprise handles the creation of a new enterprise
func (h *EnterpriseHandler) CreateEnterprise(w http.ResponseWriter, r *http.Request) {
	var enterprise domain.Enterprise
	if err := json.NewDecoder(r.Body).Decode(&enterprise); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.Service.CreateEnterprise(&enterprise); err != nil {
		http.Error(w, "Failed to create enterprise", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(enterprise)
}

// GetAllEnterprises returns a list of all enterprises
func (h *EnterpriseHandler) GetAllEnterprises(w http.ResponseWriter, r *http.Request) {
	enterpriseList, err := h.Service.GetAllEnterprises()
	if err != nil {
		http.Error(w, "Failed to retrieve enterprises", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enterpriseList)
}
