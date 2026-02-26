package model

import "time"

// Seat represents a single seat in a hall
type Seat struct {
	ID        int       `json:"id"`      // Unique seat ID
	HallID    int       `json:"hall_id"` // FK to Hall
	Row       string    `json:"row"`     // Row identifier (e.g., A, B)
	Number    string    `json:"number"`  // Seat number
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
