package models

import "fmt"

// Ticket represents a train ticket.
type Ticket struct {
	From    string  // Departure location
	To      string  // Destination
	User    User    // User who booked the ticket
	Price   float64 // Price paid for the ticket
	Seat    string  // Allocated seat number
	Section string  // Section where the seat is located
}

// NewTicket creates a new Ticket instance.
func NewTicket(from, to string, user User, pricePaid float64, seat, section string) Ticket {
	return Ticket{
		From:    from,
		To:      to,
		User:    user,
		Price:   pricePaid,
		Seat:    seat,
		Section: section,
	}
}

// String formats the ticket details.
func (t Ticket) String() string {
	return fmt.Sprintf("Ticket[From: %s, To: %s, User: %s, Price: %.2f, Seat: %s, Section: %s]",
		t.From, t.To, t.User.Email, t.Price, t.Seat, t.Section)
}
