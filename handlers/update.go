package handlers

import (
	database "CRUD_GO/Database"
	"CRUD_GO/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateItems(rw http.ResponseWriter, r *http.Request) {

	item := models.Item{}
	vars := mux.Vars(r)

	result := database.DB.First(&item, vars["id"])
	if result.Error != nil {
		http.Error(rw, "Item not found", http.StatusBadRequest)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(rw, "Decoding Error", http.StatusInternalServerError)
	}
	database.DB.Save(&item)

	json.NewEncoder(rw).Encode(item)

}
