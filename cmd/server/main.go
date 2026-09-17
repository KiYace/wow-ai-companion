package main

import (
	"fmt"
	nethttp "net/http"

	apphttp "wow-ai-companion/internal/http"
	"wow-ai-companion/internal/pet"
)

func main() {
	petService := pet.NewService()
	handler := apphttp.NewHandler(petService)

	mux := nethttp.NewServeMux()

	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /pet", handler.GetPet)
	mux.HandleFunc("POST /pet/message", handler.SendMessage)
	mux.HandleFunc("POST /pet/mood", handler.ChangeMood)

	fmt.Println("Server started on :8080")

	err := nethttp.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
