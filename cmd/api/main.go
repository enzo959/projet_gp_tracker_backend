package main

import (
	"log"
	"net/http"

	myMiddleware "github.com/enzo959/projet_gp_tracker_backend/internal/middleware"
	"github.com/joho/godotenv"

	//chi
	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/enzo959/projet_gp_tracker_backend/internal/database"
	"github.com/enzo959/projet_gp_tracker_backend/internal/handlers"
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

	// Middleware général de chi
	r.Use(chiMiddleware.Logger)
	r.Use(chiMiddleware.Recoverer)
	r.Use(myMiddleware.CORS())

	// Route publique
	r.Post("/auth/register", handlers.Register)
	r.Post("/auth/login", handlers.Login)

	// Routes sécurisées
	r.Route("/artists", func(r chi.Router) {
		r.Use(myMiddleware.JWT)
		r.Get("/", handlers.GetArtists)
		r.Get("/{id}", handlers.GetArtistByID)
		r.Get("/{id}/concerts", handlers.GetConcertsByArtist)
	})
	// routes protégées
	r.Group(func(r chi.Router) {
		r.Use(myMiddleware.JWT)

		r.Post("/concerts", handlers.CreateConcert)
		r.Get("/concerts", handlers.GetConcerts)
	})

	log.Println("Le serveur se lance sur :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
