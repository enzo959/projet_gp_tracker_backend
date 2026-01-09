package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/enzo959/projet-gp-tracker-backend/internal/database"
	"github.com/go-chi/chi/v5"
)

type Concert struct {
	ID           int       `json:"id"`
	ArtistID     int       `json:"artist_id"`
	Date         time.Time `json:"date"`
	Location     string    `json:"location"`
	PriceCents   int       `json:"price_cents"`
	TotalTickets int       `json:"total_tickets"`
}

func fetchConcerts(query string, args ...any) ([]Concert, error) {
	rows, err := database.DB.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	concerts := []Concert{}

	for rows.Next() {
		var c Concert
		if err := rows.Scan(
			&c.ID,
			&c.ArtistID,
			&c.Date,
			&c.Location,
			&c.PriceCents,
			&c.TotalTickets,
		); err != nil {
			return nil, err
		}
		concerts = append(concerts, c)
	}

	return concerts, nil
}

func GetConcerts(w http.ResponseWriter, r *http.Request) {
	concerts, err := fetchConcerts(`
		SELECT id, artist_id, date, location, price_cents, total_tickets
		FROM concerts
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(concerts)
}

func GetConcertsByArtist(w http.ResponseWriter, r *http.Request) {
	artistID := chi.URLParam(r, "id")

	concerts, err := fetchConcerts(`
		SELECT id, artist_id, date, location, price_cents, total_tickets
		FROM concerts
		WHERE artist_id = $1
	`, artistID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(concerts)
}
