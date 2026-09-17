package http

import (
	"encoding/json"
	nethttp "net/http"
)

func (h *Handler) GetPet(w nethttp.ResponseWriter, r *nethttp.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(h.petService.GetPet())
	if err != nil {
		nethttp.Error(w, "Failed to encode pet", nethttp.StatusInternalServerError)
	}
}
