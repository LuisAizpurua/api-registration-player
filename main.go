package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Message  string `json:"message"`
	Languaje string `json:"languaje"`
}

type Health struct {
	Status string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := Message{
		Message:  "Hola, mutedev",
		Languaje: "golang",
	}

	json.NewEncoder(w).Encode(response)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(Health{
		Status: "UP",
	})
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/health", status)

	fmt.Println("Servidor iniciado puerto 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
