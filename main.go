package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Pet struct {
	Name  string `json:"name"`
	Mood  string `json:"mood"`
	Level int    `json:"level"`
}

type MessageRequest struct {
	Message string `json:"message"`
}

type MessageResponse struct {
	Reply string `json:"reply"`
}

func petMessageHandler(w http.ResponseWriter, r *http.Request) {
	var request MessageRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	fmt.Println("Received message:", request.Message)

	response := MessageResponse{
		Reply: "Привет! Рад тебя видеть.",
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		fmt.Println("Failed to encode response:", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `{"status":"ok"}`)
}

func petHandler(w http.ResponseWriter, r *http.Request) {
	pet := Pet{
		Name:  "Вульпи",
		Mood:  "happy",
		Level: 1,
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(pet)
	if err != nil {
		fmt.Println("Failed to encode pet:", err)
	}
}

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/pet", petHandler)
	http.HandleFunc("/pet/message", petMessageHandler)

	fmt.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
