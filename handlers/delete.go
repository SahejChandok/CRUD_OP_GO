package handlers

import (
	database "CRUD_GO/Database"
	"CRUD_GO/models"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteItems(rw http.ResponseWriter, r *http.Request) {

	item := models.Item{}
	vars := mux.Vars(r)

	result := database.DB.First(&item, vars["id"])
	if result.Error != nil {
		http.Error(rw, "No item found", http.StatusInternalServerError)
	}

	database.DB.Delete(&item)
	rw.WriteHeader(http.StatusAccepted)

}
