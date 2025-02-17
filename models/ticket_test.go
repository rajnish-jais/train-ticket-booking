package models

import (
	"testing"
)

func TestNewTicket(t *testing.T) {
	user := NewUser("John", "Doe", "john.doe@example.com")
	ticket := NewTicket("City A", "City B", user, 50.0, "A1", "Section A")

	if ticket.From != "City A" {
		t.Errorf("Expected From to be 'City A', got %s", ticket.From)
	}
	if ticket.To != "City B" {
		t.Errorf("Expected To to be 'City B', got %s", ticket.To)
	}
	if ticket.User.Email != "john.doe@example.com" {
		t.Errorf("Expected User.Email to be 'john.doe@example.com', got %s", ticket.User.Email)
	}
	if ticket.Price != 50.0 {
		t.Errorf("Expected Price to be 50.0, got %f", ticket.Price)
	}
	if ticket.Seat != "A1" {
		t.Errorf("Expected Seat to be 'A1', got %s", ticket.Seat)
	}
	if ticket.Section != "Section A" {
		t.Errorf("Expected Section to be 'Section A', got %s", ticket.Section)
	}
}

func TestTicketString(t *testing.T) {
	user := NewUser("Alice", "Smith", "alice.smith@example.com")
	ticket := NewTicket("Town X", "Town Y", user, 75.5, "B2", "Section B")

	expected := "Ticket[From: Town X, To: Town Y, User: alice.smith@example.com, Price: 75.50, Seat: B2, Section: Section B]"
	if ticket.String() != expected {
		t.Errorf("Expected ticket string '%s', got '%s'", expected, ticket.String())
	}
}
