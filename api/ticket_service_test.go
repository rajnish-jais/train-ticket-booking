package api

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"train-ticket-booking/models"
	pb "train-ticket-booking/proto"
)

// MockStorageAdapter is a mock implementation of StorageAdapter
type MockStorageAdapter struct {
	mock.Mock
}

func (m *MockStorageAdapter) PurchaseTicket(from, to string, user models.User, pricePaid float64) (*models.Ticket, error) {
	args := m.Called(from, to, user, pricePaid)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Ticket), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockStorageAdapter) GetReceipt(email string) (*models.Receipt, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*models.Receipt), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockStorageAdapter) GetUsersBySection(section string) (*[]models.UserSeatInfo, error) {
	args := m.Called(section)
	if args.Get(0) != nil {
		return args.Get(0).(*[]models.UserSeatInfo), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockStorageAdapter) RemoveUser(email string) error {
	args := m.Called(email)
	return args.Error(0)
}

func (m *MockStorageAdapter) ModifySeat(email, newSeat, newSection string) error {
	args := m.Called(email, newSeat, newSection)
	return args.Error(0)
}

// TestPurchaseTicket validates successful and failed ticket purchases
func TestPurchaseTicket(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	service := NewTicketService(mockStorage)

	user := models.NewUser("John", "Doe", "john.doe@example.com")
	ticket := models.NewTicket("London", "Paris", user, 50.0, "A1", "A")

	mockStorage.On("PurchaseTicket", "London", "Paris", user, 50.0).Return(&ticket, nil)

	// Successful purchase
	req := &pb.PurchaseRequest{From: "London", To: "Paris", User: &pb.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"}, PricePaid: 50.0}
	resp, err := service.PurchaseTicket(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, "Ticket successfully purchased!", resp.Message)
	assert.Equal(t, "A1", resp.Ticket.Seat)
}

func TestPurchaseTicket_NoSeatsAvailable(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	service := NewTicketService(mockStorage)

	user := models.NewUser("John", "Doe", "john.doe@example.com")

	mockStorage.On("PurchaseTicket", "London", "Paris", user, 50.0).Return(nil, errors.New("no seats available"))

	// Failed purchase (No seats available)
	reqFail := &pb.PurchaseRequest{From: "London", To: "Paris", User: &pb.User{FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"}, PricePaid: 50.0}
	respFail, err := service.PurchaseTicket(context.Background(), reqFail)
	assert.Error(t, err)
	assert.Nil(t, respFail)
}

func TestGetReceipt(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	ticketService := NewTicketService(mockStorage)

	receipt := models.NewReceipt(models.NewTicket("London", "Paris", models.NewUser("John", "Doe", "john.doe@example.com"), 20.0, "A1", "A"))
	mockStorage.On("GetReceipt", "john.doe@example.com").Return(&receipt, nil)

	resp, err := ticketService.GetReceipt(context.Background(), &pb.ReceiptRequest{Email: "john.doe@example.com"})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "London", resp.Receipt.From)
	assert.Equal(t, "Paris", resp.Receipt.To)
	mockStorage.AssertExpectations(t)
}

func TestGetReceipt_NoReceiptFound(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	ticketService := NewTicketService(mockStorage)

	mockStorage.On("GetReceipt", "bob@example.com").Return(nil, errors.New("receipt not found"))

	// Failed retrieval (No receipt found)
	reqFail := &pb.ReceiptRequest{Email: "bob@example.com"}
	respFail, err := ticketService.GetReceipt(context.Background(), reqFail)
	assert.Error(t, err)
	assert.Nil(t, respFail)

	mockStorage.AssertExpectations(t)
}

func TestGetUsersBySection(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	ticketService := NewTicketService(mockStorage)

	users := []models.UserSeatInfo{{User: models.NewUser("John", "Doe", "john.doe@example.com"), Seat: "A1"}}
	mockStorage.On("GetUsersBySection", "A").Return(&users, nil)

	resp, err := ticketService.GetUsersBySection(context.Background(), &pb.UsersBySectionRequest{Section: "A"})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.UserSeatInfo, 1)
	mockStorage.AssertExpectations(t)
}

func TestGetUsersBySection_NoSectionExist(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	ticketService := NewTicketService(mockStorage)

	mockStorage.On("GetUsersBySection", "A").Return(&[]models.UserSeatInfo{}, errors.New("no users in section"))

	// Empty section
	reqEmpty := &pb.UsersBySectionRequest{Section: "A"}
	respEmpty, err := ticketService.GetUsersBySection(context.Background(), reqEmpty)
	assert.Error(t, err)
	assert.Nil(t, respEmpty)

	mockStorage.AssertExpectations(t)
}

func TestRemoveUser(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	ticketService := NewTicketService(mockStorage)

	mockStorage.On("RemoveUser", "john.doe@example.com").Return(nil)

	resp, err := ticketService.RemoveUser(context.Background(), &pb.RemoveUserRequest{Email: "john.doe@example.com"})
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	mockStorage.AssertExpectations(t)
}

func TestRemoveUser_FailedRemoval(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	ticketService := NewTicketService(mockStorage)

	mockStorage.On("RemoveUser", "bob@example.com").Return(errors.New("user not found"))

	// Failed removal
	reqFail := &pb.RemoveUserRequest{Email: "bob@example.com"}
	respFail, err := ticketService.RemoveUser(context.Background(), reqFail)
	assert.Error(t, err)
	assert.Nil(t, respFail)

	mockStorage.AssertExpectations(t)
}

func TestModifySeat(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	ticketService := NewTicketService(mockStorage)

	mockStorage.On("ModifySeat", "john.doe@example.com", "B2", "B").Return(nil)

	resp, err := ticketService.ModifySeat(context.Background(), &pb.ModifySeatRequest{Email: "john.doe@example.com", NewSeat: "B2", NewSection: "B"})
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	mockStorage.AssertExpectations(t)
}

func TestModifySeat_SeatUnavailable(t *testing.T) {
	mockStorage := new(MockStorageAdapter)
	ticketService := NewTicketService(mockStorage)

	mockStorage.On("ModifySeat", "bob@example.com", "B3", "B").Return(errors.New("seat unavailable"))

	// Failed modification
	reqFail := &pb.ModifySeatRequest{Email: "bob@example.com", NewSeat: "B3", NewSection: "B"}
	respFail, err := ticketService.ModifySeat(context.Background(), reqFail)
	assert.Error(t, err)
	assert.Nil(t, respFail)

	mockStorage.AssertExpectations(t)
}
