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

	nethttp.HandleFunc("/health", handler.Health)
	nethttp.HandleFunc("/pet", handler.GetPet)
	nethttp.HandleFunc("/pet/message", handler.SendMessage)

	fmt.Println("Server started on :8080")

	err := nethttp.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
