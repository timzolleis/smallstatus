package main

import (
	"log"
	"net/http"
	"status/api/handler"
	"status/internal/database"
)

func main() {
	database.Connect()
	database.Migrate()
	router := initializeRoutes()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server listening on port 8080...")
	server.ListenAndServe()
}

func initializeRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/monitors", handler.CreateMonitorHandler)
	mux.HandleFunc("DELETE /api/monitors/{id}", handler.DeleteMonitorHandler)
	mux.HandleFunc("GET /api/monitors/{id}", handler.GetMonitorHandler)
	return mux
}
