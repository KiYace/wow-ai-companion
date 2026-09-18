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

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeError(w, "Invalid JSON", nethttp.StatusBadRequest)
		return
	}

	request.Normalize()

	err = request.Validate()
	if err != nil {
		writeError(w, err.Error(), nethttp.StatusBadRequest)
		return
	}

	h.petService.ChangeMood(request.Mood)

	err = writeJSON(w, nethttp.StatusOK, h.petService.GetPet())
	if err != nil {
		fmt.Println("Failed to encode pet:", err)
	}
}
