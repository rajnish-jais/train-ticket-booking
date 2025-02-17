package models

import (
	"testing"
)

func TestNewReceipt(t *testing.T) {
	user := NewUser("Emily", "Davis", "emily.davis@example.com")
	ticket := NewTicket("City C", "City D", user, 100.0, "C3", "Section C")
	receipt := NewReceipt(ticket)

	if receipt.From != "City C" {
		t.Errorf("Expected From to be 'City C', got %s", receipt.From)
	}
	if receipt.To != "City D" {
		t.Errorf("Expected To to be 'City D', got %s", receipt.To)
	}
	if receipt.User.Email != "emily.davis@example.com" {
		t.Errorf("Expected User.Email to be 'emily.davis@example.com', got %s", receipt.User.Email)
	}
	if receipt.PricePaid != 100.0 {
		t.Errorf("Expected PricePaid to be 100.0, got %f", receipt.PricePaid)
	}
}

func TestReceiptString(t *testing.T) {
	user := NewUser("Chris", "Johnson", "chris.johnson@example.com")
	ticket := NewTicket("Town M", "Town N", user, 150.0, "D4", "Section D")
	receipt := NewReceipt(ticket)

	expected := "Receipt[From: Town M, To: Town N, User: chris.johnson@example.com, Price Paid: 150.00]"
	if receipt.String() != expected {
		t.Errorf("Expected receipt string '%s', got '%s'", expected, receipt.String())
	}
}
