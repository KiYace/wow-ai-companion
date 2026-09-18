package http

import (
	"fmt"
	nethttp "net/http"
)

func (h *Handler) GetPet(w nethttp.ResponseWriter, r *nethttp.Request) {
	err := writeJSON(w, nethttp.StatusOK, h.petService.GetPet())
	if err != nil {
		fmt.Println("Failed to encode pet:", err)
	}
}
