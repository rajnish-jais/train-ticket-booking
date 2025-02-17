package models

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	user := NewUser("Jane", "Doe", "jane.doe@example.com")

	if user.FirstName != "Jane" {
		t.Errorf("Expected FirstName to be 'Jane', got %s", user.FirstName)
	}
	if user.LastName != "Doe" {
		t.Errorf("Expected LastName to be 'Doe', got %s", user.LastName)
	}
	if user.Email != "jane.doe@example.com" {
		t.Errorf("Expected Email to be 'jane.doe@example.com', got %s", user.Email)
	}
}

func TestUserString(t *testing.T) {
	user := NewUser("Mike", "Brown", "mike.brown@example.com")
	expected := "User[Name: Mike Brown, Email: mike.brown@example.com]"

	if user.String() != expected {
		t.Errorf("Expected user string '%s', got '%s'", expected, user.String())
	}
}
