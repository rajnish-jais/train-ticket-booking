package models

import "fmt"

// Receipt represents a payment receipt for a ticket purchase.
type Receipt struct {
	From      string
	To        string
	User      User
	PricePaid float64
}

// NewReceipt creates a receipt for a purchased ticket.
func NewReceipt(ticket Ticket) Receipt {
	return Receipt{
		From:      ticket.From,
		To:        ticket.To,
		User:      ticket.User,
		PricePaid: ticket.Price,
	}
}

// String formats the receipt details.
func (r Receipt) String() string {
	return fmt.Sprintf("Receipt[From: %s, To: %s, User: %s, Price Paid: %.2f]",
		r.From, r.To, r.User.Email, r.PricePaid)
}
