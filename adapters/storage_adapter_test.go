package adapters

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"train-ticket-booking/models"
)

func setupInMemoryAdapter() *InMemoryAdapter {
	// Initialize the in-memory adapter for each test
	return NewInMemoryAdapter()
}

func TestPurchaseTicket(t *testing.T) {
	// Initialize the in-memory adapter
	adapter := setupInMemoryAdapter()

	// Create a user
	user := models.NewUser("John", "Doe", "john.doe@example.com")

	// Purchase a ticket
	ticket, err := adapter.PurchaseTicket("London", "Paris", user, 20.0)
	assert.NoError(t, err)
	assert.NotNil(t, ticket)
	assert.Equal(t, "London", ticket.From)
	assert.Equal(t, "Paris", ticket.To)
	assert.Equal(t, user.Email, ticket.User.Email)
}

func TestPurchaseTicket_NoSeatsAvailable(t *testing.T) {
	// Initialize the in-memory adapter
	adapter := setupInMemoryAdapter()

	// Create users and purchase tickets until all seats are occupied
	user1 := models.NewUser("John", "Doe", "john.doe@example.com")
	user2 := models.NewUser("Jane", "Smith", "jane.smith@example.com")
	user3 := models.NewUser("Alice", "Johnson", "alice.johnson@example.com")
	user4 := models.NewUser("Bob", "Brown", "bob.brown@example.com")
	user5 := models.NewUser("Charlie", "Davis", "charlie.davis@example.com")
	user6 := models.NewUser("David", "Wilson", "david.wilson@example.com")

	// Purchase tickets for 6 users (this should fill all available seats)
	_, err := adapter.PurchaseTicket("London", "Paris", user1, 20.0)
	assert.NoError(t, err)
	_, err = adapter.PurchaseTicket("London", "Paris", user2, 20.0)
	assert.NoError(t, err)
	_, err = adapter.PurchaseTicket("London", "Paris", user3, 20.0)
	assert.NoError(t, err)
	_, err = adapter.PurchaseTicket("London", "Paris", user4, 20.0)
	assert.NoError(t, err)
	_, err = adapter.PurchaseTicket("London", "Paris", user5, 20.0)
	assert.NoError(t, err)
	_, err = adapter.PurchaseTicket("London", "Paris", user6, 20.0)
	assert.NoError(t, err)

	// Attempt to purchase a 7th ticket (this should fail as no seats are available)
	_, err = adapter.PurchaseTicket("London", "Berlin", user1, 25.0)
	assert.Error(t, err)
	assert.Equal(t, "no available seats", err.Error())
}

func TestGetReceipt(t *testing.T) {
	// Initialize the in-memory adapter
	adapter := setupInMemoryAdapter()

	// Create a user and purchase a ticket
	user := models.NewUser("John", "Doe", "john.doe@example.com")
	_, err := adapter.PurchaseTicket("London", "Paris", user, 20.0)
	assert.NoError(t, err)

	// Get the receipt for the user
	receipt, err := adapter.GetReceipt(user.Email)
	assert.NoError(t, err)
	assert.NotNil(t, receipt)
	assert.Equal(t, user.Email, receipt.User.Email)
	assert.Equal(t, 20.0, receipt.PricePaid)
}

func TestGetReceipt_UserNotFound(t *testing.T) {
	// Initialize the in-memory adapter
	adapter := setupInMemoryAdapter()

	// Try to get a receipt for a non-existing user
	receipt, err := adapter.GetReceipt("non.existing.user@example.com")
	assert.Error(t, err)
	assert.Nil(t, receipt)
	assert.Equal(t, "receipt not found", err.Error())
}

func TestGetUsersBySection(t *testing.T) {
	// Initialize the in-memory adapter
	adapter := setupInMemoryAdapter()

	// Create users and purchase tickets
	user1 := models.NewUser("John", "Doe", "john.doe@example.com")
	user2 := models.NewUser("Jane", "Smith", "jane.smith@example.com")
	_, err := adapter.PurchaseTicket("London", "Paris", user1, 20.0)
	assert.NoError(t, err)
	ticket, err := adapter.PurchaseTicket("London", "Paris", user2, 20.0)
	assert.NoError(t, err)

	_, err = adapter.GetUsersBySection(ticket.Section)
	assert.NoError(t, err)

}

func TestGetUsersBySection_SectionEmpty(t *testing.T) {
	// Initialize the in-memory adapter
	adapter := setupInMemoryAdapter()

	// Try to get users for an empty section
	users, err := adapter.GetUsersBySection("B")
	assert.Error(t, err) // Expecting an error since section B is empty
	assert.Nil(t, users) // Ensure users is nil when section is empty
}

func TestRemoveUser(t *testing.T) {
	// Initialize the in-memory adapter
	adapter := setupInMemoryAdapter()

	// Case 1: Attempt to remove a user that does not exist
	err := adapter.RemoveUser("unknown@example.com")
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	// Case 2: Purchase a ticket for a user
	user := models.NewUser("John", "Doe", "john.doe@example.com")
	ticket, err := adapter.PurchaseTicket("London", "Paris", user, 20.0)
	assert.NoError(t, err)

	// Capture assigned seat & section
	seat := ticket.Seat
	section := ticket.Section

	// Verify the seat is occupied
	assert.True(t, adapter.sections[section][seat], "Seat should be occupied before removal")

	// Case 3: Successfully remove the user
	err = adapter.RemoveUser(user.Email)
	assert.NoError(t, err)

	// Case 4: Ensure the ticket is deleted
	_, exists := adapter.tickets[user.Email]
	assert.False(t, exists, "User's ticket should be removed from storage")

	// Case 5: Ensure the seat is now available
	assert.False(t, adapter.sections[section][seat], "Seat should be available after user removal")

	// Case 6: Try to remove the same user again (should return "user not found")
	err = adapter.RemoveUser(user.Email)
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())
}

func TestModifySeat(t *testing.T) {
	// Initialize the in-memory adapter
	adapter := setupInMemoryAdapter()

	// Case 1: Modify seat for a user who doesn't exist
	err := adapter.ModifySeat("unknown@example.com", "A2", "A")
	assert.Error(t, err)
	assert.Equal(t, "user not found", err.Error())

	// Case 2: Purchase a ticket for a user
	user := models.NewUser("John", "Doe", "john.doe@example.com")
	ticket, err := adapter.PurchaseTicket("London", "Paris", user, 20.0)
	assert.NoError(t, err)

	// Capture the original seat
	originalSeat := ticket.Seat
	originalSection := ticket.Section

	// Case 3: Attempt to modify the seat to an occupied one (should fail)
	err = adapter.ModifySeat(user.Email, originalSeat, originalSection) // Trying to "move" to the same seat
	assert.Error(t, err)
	assert.Equal(t, "new seat is not available", err.Error())

	// Case 4: Modify the seat to another available seat (should pass)
	for sec, seats := range adapter.sections {
		for seat, occupied := range seats {
			if !occupied {
				err = adapter.ModifySeat(user.Email, seat, sec)
				assert.NoError(t, err)

				// Ensure the seat changed
				assert.NotEqual(t, originalSeat, ticket.Seat, "Seat should be updated after modification")
				return // Exit test after first successful seat change
			}
		}
	}
}
