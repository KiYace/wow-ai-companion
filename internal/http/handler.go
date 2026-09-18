package http

import (
	"encoding/json"
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

func writeJSON(w nethttp.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}

func writeError(w nethttp.ResponseWriter, message string, status int) {
	nethttp.Error(w, message, status)
}
