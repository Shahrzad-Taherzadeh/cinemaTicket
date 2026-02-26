package model

import (
	"time"
)

// Ticket represents a booked seat
type Ticket struct {
	ID          int        `json:"id"`           // Unique ticket ID
	UserID      int        `json:"user_id"`      // FK to User
	ScreeningID int        `json:"screening_id"` // FK to Screening
	ShowTimeID  int        `json:"show_time_id"` // FK to ShowTime
	ScreenSeat  ScreenSeat `json:"screen_seat"`  // Reserved or sold seat info
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
