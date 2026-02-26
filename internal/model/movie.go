package model

import "time"

// Movie represents a film
type Movie struct {
	ID              int       `json:"id"`               // Unique movie ID
	Name            string    `json:"name"`             // Movie title
	Year            int       `json:"year"`             // Release year
	Genre           string    `json:"genre"`            // Genre name
	Director        string    `json:"director"`         // Director name
	DurationMinutes int       `json:"duration_minutes"` // Duration in minutes
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

// Duration returns the total duration of the session
func (m *Movie) Duration() time.Duration {
	return time.Duration(m.DurationMinutes) * time.Minute
}
