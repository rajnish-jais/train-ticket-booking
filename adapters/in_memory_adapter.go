package adapters

import (
	"errors"
	"sync"
	"train-ticket-booking/models"
)

// InMemoryAdapter is an in-memory storage with dynamic seat allocation.
type InMemoryAdapter struct {
	mu       sync.Mutex
	tickets  map[string]*models.Ticket
	sections map[string]map[string]bool // Tracks seat availability (true = occupied, false = available)
}

// NewInMemoryAdapter initializes storage with dynamic seat allocation.
func NewInMemoryAdapter() *InMemoryAdapter {
	return &InMemoryAdapter{
		tickets: make(map[string]*models.Ticket),
		sections: map[string]map[string]bool{
			"A": {"A1": false, "A2": false, "A3": false}, // All seats available initially
			"B": {"B1": false, "B2": false, "B3": false},
		},
	}
}

// PurchaseTicket reserves a seat and creates a ticket.
func (s *InMemoryAdapter) PurchaseTicket(from, to string, user models.User, pricePaid float64) (*models.Ticket, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Find the first available seat
	var assignedSeat, assignedSection string
	for section, seats := range s.sections {
		for seat, occupied := range seats {
			if !occupied {
				assignedSeat = seat
				assignedSection = section
				s.sections[section][seat] = true // Mark seat as occupied
				break
			}
		}
		if assignedSeat != "" {
			break
		}
	}

	if assignedSeat == "" {
		return nil, errors.New("no available seats")
	}

	// Create and store the ticket
	ticket := models.NewTicket(from, to, user, pricePaid, assignedSeat, assignedSection)
	s.tickets[user.Email] = &ticket
	return &ticket, nil
}

// GetReceipt retrieves the ticket receipt by email.
func (s *InMemoryAdapter) GetReceipt(email string) (*models.Receipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ticket, exists := s.tickets[email]
	if !exists {
		return nil, errors.New("receipt not found")
	}

	receipt := models.NewReceipt(*ticket)
	return &receipt, nil
}

// GetUsersBySection retrieves all users in a given section.
func (s *InMemoryAdapter) GetUsersBySection(section string) (*[]models.UserSeatInfo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var users []models.UserSeatInfo
	for _, ticket := range s.tickets {
		if ticket.Section == section {
			users = append(users, models.UserSeatInfo{
				User: ticket.User,
				Seat: ticket.Seat,
			})
		}
	}

	if len(users) == 0 {
		return nil, errors.New("no users found in this section")
	}
	return &users, nil
}

// RemoveUser removes a user and frees up their seat.
func (s *InMemoryAdapter) RemoveUser(email string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	ticket, exists := s.tickets[email]
	if !exists {
		return errors.New("user not found")
	}

	// Free the seat
	s.sections[ticket.Section][ticket.Seat] = false
	delete(s.tickets, email)

	return nil
}

// ModifySeat changes a user's seat and updates availability.
func (s *InMemoryAdapter) ModifySeat(email, newSeat, newSection string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	ticket, exists := s.tickets[email]
	if !exists {
		return errors.New("user not found")
	}

	// Check if new seat is available
	if occupied, ok := s.sections[newSection][newSeat]; !ok || occupied {
		return errors.New("new seat is not available")
	}

	// Free the old seat
	s.sections[ticket.Section][ticket.Seat] = false

	// Assign the new seat
	ticket.Seat = newSeat
	ticket.Section = newSection
	s.sections[newSection][newSeat] = true

	return nil
}
