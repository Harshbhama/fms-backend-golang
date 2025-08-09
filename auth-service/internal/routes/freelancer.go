package routes

import (
	"github.com/yourusername/auth-service/internal/handlers"
	"net/http"
)

func RegisterFreelancerRoutes(mux *http.ServeMux, handler *handlers.FreelancerHandler) {
	mux.HandleFunc("/freelancers", handler.CreateFreelancer)
}
