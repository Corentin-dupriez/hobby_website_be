package recipes

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListRecipes(w http.ResponseWriter, r *http.Request) {
	// call the service -> ListRecipes
	// Return JSON in an HTTP response
	recipes := []string{"Hello", "World"}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(recipes)
	if err != nil {
		slog.Error("Error encoding products")
	}
}
