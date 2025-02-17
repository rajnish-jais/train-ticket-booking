package api

import (
	"context"
	"fmt"

	"train-ticket-booking/adapters"
	"train-ticket-booking/models"
	pb "train-ticket-booking/proto"
)

// TicketService implements the gRPC TrainTicketService
type TicketService struct {
	pb.UnimplementedTrainTicketServiceServer
	storage adapters.StorageAdapter
}

// NewTicketService initializes a new TicketService
func NewTicketService(storage adapters.StorageAdapter) *TicketService {
	return &TicketService{storage: storage}
}

// PurchaseTicket handles the purchase of a ticket
func (s *TicketService) PurchaseTicket(ctx context.Context, req *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	user := models.NewUser(req.User.FirstName, req.User.LastName, req.User.Email)
	ticket, err := s.storage.PurchaseTicket(req.From, req.To, user, req.PricePaid)
	if err != nil {
		return nil, fmt.Errorf("failed to purchase ticket: %w", err)
	}

	return &pb.PurchaseResponse{
		Ticket: &pb.Ticket{
			From:    ticket.From,
			To:      ticket.To,
			User:    &pb.User{FirstName: ticket.User.FirstName, LastName: ticket.User.LastName, Email: ticket.User.Email},
			Price:   ticket.Price,
			Seat:    ticket.Seat,
			Section: ticket.Section,
		},
		Message: "Ticket successfully purchased!",
	}, nil
}

// GetReceipt fetches a receipt for a user based on email
func (s *TicketService) GetReceipt(ctx context.Context, req *pb.ReceiptRequest) (*pb.ReceiptResponse, error) {
	receipt, err := s.storage.GetReceipt(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve receipt: %w", err)
	}

	return &pb.ReceiptResponse{
		Receipt: &pb.Receipt{
			From:      receipt.From,
			To:        receipt.To,
			User:      &pb.User{FirstName: receipt.User.FirstName, LastName: receipt.User.LastName, Email: receipt.User.Email},
			PricePaid: receipt.PricePaid,
		},
		Message: "Receipt retrieved successfully",
	}, nil
}

// GetUsersBySection fetches all users in a given section
func (s *TicketService) GetUsersBySection(ctx context.Context, req *pb.UsersBySectionRequest) (*pb.UsersBySectionResponse, error) {
	userSeatInfos, err := s.storage.GetUsersBySection(req.Section)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve users in section: %w", err)
	}

	var userSeats []*pb.UserSeatInfo
	for _, info := range *userSeatInfos {
		userSeats = append(userSeats, &pb.UserSeatInfo{
			User: &pb.User{
				FirstName: info.User.FirstName,
				LastName:  info.User.LastName,
				Email:     info.User.Email,
			},
			Seat: info.Seat,
		})
	}

	return &pb.UsersBySectionResponse{
		UserSeatInfo: userSeats,
		Message:      fmt.Sprintf("Users in section %s retrieved successfully", req.Section),
	}, nil
}

// RemoveUser removes a user from the train based on email
func (s *TicketService) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	err := s.storage.RemoveUser(req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to remove user: %w", err)
	}

	return &pb.RemoveUserResponse{
		Success: true,
		Message: "User successfully removed",
	}, nil
}

// ModifySeat changes a user's seat
func (s *TicketService) ModifySeat(ctx context.Context, req *pb.ModifySeatRequest) (*pb.ModifySeatResponse, error) {
	err := s.storage.ModifySeat(req.Email, req.NewSeat, req.NewSection)
	if err != nil {
		return nil, fmt.Errorf("failed to modify seat: %w", err)
	}

	return &pb.ModifySeatResponse{
		Success:    true,
		NewSeat:    req.NewSeat,
		NewSection: req.NewSection,
		Message:    "Seat successfully modified",
	}, nil
}
