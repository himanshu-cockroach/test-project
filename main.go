package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WelcomeMessage struct {
	Message     string `json:"message"`
	CurrentTime string `json:"current_time"`
}

func main() {
	// defining router
	mux := http.NewServeMux()
	mux.HandleFunc("/landing", helloWorld)

	// starting server
	fmt.Println("Server is running at 127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("function called")
	currentTime := []WelcomeMessage{{
		Message:     "Hello World",
		CurrentTime: http.TimeFormat,
	}}

	json.NewEncoder(w).Encode(currentTime)
}
