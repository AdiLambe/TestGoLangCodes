package handler

import (
	"encoding/json"
	"my_project/app/dto"
	"my_project/app/service"
	"net/http"
)

type EnterpriseHandler struct {
	Service *service.EnterpriseService
}

func (h *EnterpriseHandler) CreateEnterprise(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateEnterpriseRequest
	json.NewDecoder(r.Body).Decode(&req)

	enterprise, err := h.Service.CreateEnterprise(req.Name, req.Region)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(enterprise)
}
