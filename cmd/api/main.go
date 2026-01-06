package main

import (
	"encoding/json"
	"log"
	"net/http"

	//chi
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"statu": "ok",
	})
}

func main() {
	r := chi.NewRouter()

	// middlewares de base
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", healthHandler)

	log.Println("Le serveur se lance sur :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
