package handlers

import (
	"github.com/yourusername/auth-service/internal/models"
	"github.com/yourusername/auth-service/internal/services"
	"encoding/json"
	"net/http"
)

type FreelancerHandler struct {
	Service *services.FreelancerService
}

func NewFreelancerHandler(service *services.FreelancerService) *FreelancerHandler {
	return &FreelancerHandler{Service: service}
}

func (h *FreelancerHandler) CreateFreelancer(w http.ResponseWriter, r *http.Request) {
	var freelancer models.Freelancer
	if err := json.NewDecoder(r.Body).Decode(&freelancer); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Service.CreateFreelancer(&freelancer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(freelancer)
}
