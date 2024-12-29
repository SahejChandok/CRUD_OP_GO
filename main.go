package main

import (
	database "CRUD_GO/Database"
	"CRUD_GO/handlers"
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDB()

	sm := mux.NewRouter()

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/items", handlers.CreateItem)

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/items", handlers.GetItem)

	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/items/{id:[0-9]+}", handlers.DeleteItems)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	s.ListenAndServe()

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
