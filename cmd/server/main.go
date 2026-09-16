package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wow-ai-companion/internal/pet"
)

func main() {
	petService := pet.NewService()

	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/pet", petHandler(petService))
	http.HandleFunc("/pet/message", petMessageHandler(petService))

	fmt.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok"}`)
}

func petHandler(petService *pet.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(petService.GetPet())
		if err != nil {
			fmt.Println("Failed to encode pet:", err)
		}
	}
}

func petMessageHandler(petService *pet.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var request struct {
			Message string `json:"message"`
		}

		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		response := struct {
			Reply string `json:"reply"`
		}{
			Reply: petService.RespondTo(request.Message),
		}

		w.Header().Set("Content-Type", "application/json")

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			fmt.Println("Failed to encode response:", err)
		}
	}
}
