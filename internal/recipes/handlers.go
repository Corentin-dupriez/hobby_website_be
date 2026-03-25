package recipes

import "net/http"

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
}
