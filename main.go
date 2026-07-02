package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Message  string `json:"message"`
	Languaje string `json:"languaje"`
}

type Health struct {
	Status string `json:"status"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	w.Header().Set("Content-Type", "application/json")

	response := Message{
		Message:  "Hola, " + name,
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

func showEnv() string {
	nodeIP := os.Getenv("NODE_IP")
	svcPort := os.Getenv("SVC_PORT")

	return fmt.Sprintf("http://%s:%s", nodeIP, svcPort)
}

func main() {
	http.HandleFunc("/hello", handler)
	http.HandleFunc("/health", status)

	fmt.Println("Contenedor iniciado en el puerto 8080")
	fmt.Println(showEnv())

	log.Fatal(http.ListenAndServe(":8080", nil))
}
