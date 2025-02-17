package adapters

import "train-ticket-booking/models"

// StorageAdapter defines the interface for ticket storage operations.
type StorageAdapter interface {
	PurchaseTicket(from, to string, user models.User, pricePaid float64) (*models.Ticket, error)
	GetReceipt(email string) (*models.Receipt, error)
	GetUsersBySection(section string) (*[]models.UserSeatInfo, error)
	RemoveUser(email string) error
	ModifySeat(email, newSeat, newSection string) error
}
