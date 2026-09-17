package http

import (
	"fmt"
	nethttp "net/http"

	"wow-ai-companion/internal/pet"
)

type Handler struct {
	petService *pet.Service
}

func NewHandler(petService *pet.Service) *Handler {
	return &Handler{
		petService: petService,
	}
}

func (h *Handler) Health(w nethttp.ResponseWriter, r *nethttp.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok"}`)
}
