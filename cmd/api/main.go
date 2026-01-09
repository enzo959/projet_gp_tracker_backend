package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"

	//chi
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/enzo959/projet-gp-tracker-backend/internal/database"
	"github.com/enzo959/projet-gp-tracker-backend/internal/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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
	r.Get("/artists/{id}", handlers.GetArtistByID)
	r.Get("/artists/{id}/concerts", handlers.GetConcertsByArtist)
	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/login", handlers.Login)

	log.Println("Le serveur se lance sur :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
