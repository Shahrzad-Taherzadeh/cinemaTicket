package model

import "time"

// ShowTime represents a movie session
type ShowTime struct {
	ID        int          `json:"id"`      // Unique ShowTime ID
	HallID    int          `json:"hall_id"` // FK to Hall
	Date      ShowTimeDate `json:"date"`    // Session start and end times
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type ShowTimeDate struct {
	Start time.Time `json:"start"` // Session start time
	End   time.Time `json:"end"`   // Session end time
}

// Duration returns the total duration of the session
func (s *ShowTime) Duration() time.Duration {
	return s.Date.End.Sub(s.Date.Start)
}
