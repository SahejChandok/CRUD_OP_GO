package handlers

import (
	database "CRUD_GO/Database"
	"CRUD_GO/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetItem(rw http.ResponseWriter, r *http.Request) {

	//to return slice, if you want to get all the users
	//item := []models.Item{}

	item := models.Item{}

	vars := mux.Vars(r)

	result := database.DB.First(&item, vars["id"])
	//result := database.DB.Find(&items)
	if result.Error != nil {
		http.Error(rw, "Error insertin in db", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(item)
}
