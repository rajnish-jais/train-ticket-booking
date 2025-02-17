package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"train-ticket-booking/config" // Import config package
	pb "train-ticket-booking/proto"

	"google.golang.org/grpc"
)

// TrainTicketClient struct to hold the gRPC connection
type TrainTicketClient struct {
	conn   *grpc.ClientConn
	client pb.TrainTicketServiceClient
}

// NewTrainTicketClient initializes a new client connection
func NewTrainTicketClient() (*TrainTicketClient, error) {
	// Load configuration internally
	cfg := config.MustLoadConfig("config/config.yaml")
	serverAddr := fmt.Sprintf("localhost:%s", cfg.Server.Port)

	// Establish gRPC connection
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	client := pb.NewTrainTicketServiceClient(conn)
	return &TrainTicketClient{conn: conn, client: client}, nil
}

// CloseConnection closes the gRPC connection
func (c *TrainTicketClient) CloseConnection() {
	c.conn.Close()
}

// PurchaseTicket calls the PurchaseTicket gRPC method
func (c *TrainTicketClient) PurchaseTicket(from, to, firstName, lastName, email string, pricePaid float64) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.PurchaseRequest{
		From: from,
		To:   to,
		User: &pb.User{
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
		},
		PricePaid: pricePaid,
	}

	resp, err := c.client.PurchaseTicket(ctx, req)
	if err != nil {
		log.Fatalf("Could not purchase ticket: %v", err)
	}

	fmt.Printf("Purchase successful: %s\n", resp.Message)
	fmt.Printf("Ticket details: %+v\n", resp.Ticket)
}

// GetReceipt calls the GetReceipt gRPC method
func (c *TrainTicketClient) GetReceipt(email string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.ReceiptRequest{Email: email}
	resp, err := c.client.GetReceipt(ctx, req)
	if err != nil {
		log.Fatalf("Could not retrieve receipt: %v", err)
	}

	fmt.Printf("Receipt retrieved successfully: %s\n", resp.Message)
	fmt.Printf("Receipt details: %+v\n", resp.Receipt)
}

// GetUsersBySection calls the GetUsersBySection gRPC method
func (c *TrainTicketClient) GetUsersBySection(section string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.UsersBySectionRequest{Section: section}
	resp, err := c.client.GetUsersBySection(ctx, req)
	if err != nil {
		log.Fatalf("Could not retrieve users by section: %v", err)
	}

	fmt.Printf("Users in section %s: %s\n", section, resp.Message)
	for _, userSeat := range resp.UserSeatInfo {
		fmt.Printf("User: %s %s, Email: %s, Seat: %s\n",
			userSeat.User.FirstName, userSeat.User.LastName, userSeat.User.Email, userSeat.Seat)
	}
}

// RemoveUser calls the RemoveUser gRPC method
func (c *TrainTicketClient) RemoveUser(email string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.RemoveUserRequest{Email: email}
	resp, err := c.client.RemoveUser(ctx, req)
	if err != nil {
		log.Fatalf("Could not remove user: %v", err)
	}

	fmt.Printf("User removed: %s\n", resp.Message)
}

// ModifySeat calls the ModifySeat gRPC method
func (c *TrainTicketClient) ModifySeat(email, newSeat, newSection string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &pb.ModifySeatRequest{
		Email:      email,
		NewSeat:    newSeat,
		NewSection: newSection,
	}

	resp, err := c.client.ModifySeat(ctx, req)
	if err != nil {
		fmt.Printf("Could not modify seat: %v", err)
		return
	}

	fmt.Printf("Seat modified: %s\n", resp.Message)
	fmt.Printf("New Seat: %s, New Section: %s\n", resp.NewSeat, resp.NewSection)
}
