package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Text string `json:"text"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Text: "Welcome to the API!",
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/", homeHandler)

	http.ListenAndServe(":8080", nil)
}
