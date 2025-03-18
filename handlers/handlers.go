package handlers

import (
	"encoding/json"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"message": "Welcome to the Home Page",
		"author":  "Aaron Mineen",
		"version": "v0",
	}
	json.NewEncoder(w).Encode(data)
}
