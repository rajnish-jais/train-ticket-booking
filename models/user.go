package models

import "fmt"

// User represents a passenger booking a ticket.
type User struct {
	FirstName string // First name of the user
	LastName  string // Last name of the user
	Email     string // Email identifier for the user
}

// NewUser creates a new User instance.
func NewUser(firstName, lastName, email string) User {
	return User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}

// String formats user details.
func (u User) String() string {
	return fmt.Sprintf("User[Name: %s %s, Email: %s]", u.FirstName, u.LastName, u.Email)
}
