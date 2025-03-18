package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/amineen/utility-api/services"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := map[string]string{
		"message": "Welcome to the Home Page",
		"author":  "Aaron Mineen",
		"version": "v0.0.1",
	}
	json.NewEncoder(w).Encode(data)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := services.GetAllCustomers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func GetCustomersByUtilityID(w http.ResponseWriter, r *http.Request) {
	utilityID := r.PathValue("utilityId")
	customers, err := services.GetCustomersByUtilityID(utilityID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
