package handlers

import (
	database "CRUD_GO/Database"
	"CRUD_GO/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetItem(rw http.ResponseWriter, r *http.Request){

	item := models.Item{}
	
	vars := mux.Vars(r)

	result:= database.DB.First(&item, vars["id"])
	if(result.Error !=nil){
		http.Error(rw, "Error insertin in db", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(rw).Encode(item)
}