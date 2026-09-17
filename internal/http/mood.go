package http

import (
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"strings"
)

type ChangeMoodRequest struct {
	Mood string `json:"mood"`
}

func (r ChangeMoodRequest) Validate() error {
	if r.Mood == "" {
		return fmt.Errorf("mood is required")
	}

	return nil
}

func (r *ChangeMoodRequest) Normalize() {
	r.Mood = strings.TrimSpace(r.Mood)
}

func (h *Handler) ChangeMood(w nethttp.ResponseWriter, r *nethttp.Request) {
	var request ChangeMoodRequest

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		nethttp.Error(w, "Invalid JSON", nethttp.StatusBadRequest)
		return
	}

	request.Normalize()

	err = request.Validate()
	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusBadRequest)
		return
	}

	h.petService.ChangeMood(request.Mood)

	err = json.NewEncoder(w).Encode(h.petService.GetPet())
	if err != nil {
		nethttp.Error(w, "Failed to encode pet", nethttp.StatusInternalServerError)
	}
}
