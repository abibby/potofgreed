package main

import (
	"encoding/json"
	"net/http"

	"github.com/zwzn/potofgreed/_example/models"
)

//go:generate go run ../potofgreed/main.go

func greet(w http.ResponseWriter, r *http.Request) {
	b := models.Book{}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}
