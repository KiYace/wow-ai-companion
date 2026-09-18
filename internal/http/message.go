package http

import (
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"strings"
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

func (r *SendMessageRequest) Normalize() {
	r.Message = strings.TrimSpace(r.Message)
}

func (h *Handler) SendMessage(w nethttp.ResponseWriter, r *nethttp.Request) {
	var request SendMessageRequest

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

	response := SendMessageResponse{
		Reply: h.petService.RespondTo(request.Message),
	}

	err = writeJSON(w, nethttp.StatusOK, response)
	if err != nil {
		fmt.Println("Failed to encode response:", err)
	}
}
