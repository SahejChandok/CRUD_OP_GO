package handlers

import (
	database "CRUD_GO/Database"
	"CRUD_GO/models"
	"encoding/json"
	"net/http"
)

func CreateItem(rw http.ResponseWriter, r *http.Request) {
	item := models.Item{}

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(rw, "Decoding error", http.StatusInternalServerError)
		return
	}
	result := database.DB.Create(&item)
	if result.Error != nil {
		http.Error(rw, "Connection Error", http.StatusInternalServerError)
		return
	}

	//rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(item)
}
