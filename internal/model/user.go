package model

// User represents a system user with minimal info
type User struct {
	ID    int    `json:"id"`    // Unique user ID
	Name  string `json:"name"`  // User's full name
	Phone string `json:"phone"` // User's phone number
}
