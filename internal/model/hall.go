package model

import "time"

// Hall represents a cinema hall
type Hall struct {
	ID        int       `json:"id"`     // Unique hall ID
	Name      string    `json:"name"`   // Hall name
	Number    string    `json:"number"` // Hall number
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
