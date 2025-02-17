package models

// UserSeatInfo represents a user's seat allocation details.
type UserSeatInfo struct {
	User User   // Embedded User struct
	Seat string // Allocated seat number
}
