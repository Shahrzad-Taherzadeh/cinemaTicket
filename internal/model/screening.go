package model

import "time"

// Screening represents a movie screening
type Screening struct {
	ID             int            `json:"id"`              // Unique screening ID
	ShowTimeID     int            `json:"show_time_id"`    // FK to ShowTime
	MovieID        int            `json:"movie_id"`        // FK to Movie
	ScreeningSeats ScreeningSeats `json:"screening_seats"` // Seats reserved and sold
	Price          int64          `json:"price"`           // Ticket price
	IsActive       bool           `json:"is_active"`       // Screening status
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type ScreeningSeats struct {
	Reserved []ScreenSeat `json:"reserved"` // Seats that are reserved
	Sold     []ScreenSeat `json:"sold"`     // Seats that are sold
}

type ScreenSeat struct {
	SeatID    int       `json:"seat_id"`    // FK to Seat
	Row       string    `json:"row"`        // Seat row
	Number    string    `json:"column"`     // Seat Number
	IssueTime time.Time `json:"issue_time"` // When seat was reserved or sold
}

// Returns number of sold seats
func (s *Screening) SoldCount() int {
	return len(s.ScreeningSeats.Sold)
}

// Returns number of reserved seats
func (s *Screening) ReservedCount() int {
	return len(s.ScreeningSeats.Reserved)
}
