package http

import (
	"encoding/json"
	"fmt"
	nethttp "net/http"
)

type SendMessageRequest struct {
	Message string `json:"message"`
}

type SendMessageResponse struct {
	Reply string `json:"reply"`
}

func (r SendMessageRequest) Validate() error {
	if r.Message == "" {
		return fmt.Errorf("message is required")
	}

	return nil
}

func (h *Handler) SendMessage(w nethttp.ResponseWriter, r *nethttp.Request) {
	var request SendMessageRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		nethttp.Error(w, "Invalid JSON", nethttp.StatusBadRequest)
		return
	}

	err = request.Validate()
	if err != nil {
		nethttp.Error(w, err.Error(), nethttp.StatusBadRequest)
		return
	}

	response := SendMessageResponse{
		Reply: h.petService.RespondTo(request.Message),
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("Failed to encode response:", err)
	}
}
