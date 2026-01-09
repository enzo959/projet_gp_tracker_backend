package main

import (
	"log"
	"net/http"

	//chi
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/enzo959/projet-gp-tracker-backend/internal/database"
	"github.com/enzo959/projet-gp-tracker-backend/internal/handlers"
)

func main() {
	// connexion postgreSQL
	if err := database.Connect(); err != nil {
		log.Fatal("DB connection failed:", err)
	}
	// routing
	r := chi.NewRouter()

	// middlewares de base fournit par chi
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", handlers.GetHealth)
	r.Get("/artists", handlers.GetArtists)
	r.Get("/concerts", handlers.GetConcerts)

	log.Println("Le serveur se lance sur :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
