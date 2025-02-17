package client

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	pb "train-ticket-booking/proto"
)

// MockServer is a mock implementation of the gRPC TrainTicketServiceServer
type MockServer struct {
	pb.UnimplementedTrainTicketServiceServer
	mock.Mock
}

// Mock PurchaseTicket implementation
func (m *MockServer) PurchaseTicket(ctx context.Context, req *pb.PurchaseRequest) (*pb.PurchaseResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.PurchaseResponse), args.Error(1)
}

// Mock GetReceipt implementation
func (m *MockServer) GetReceipt(ctx context.Context, req *pb.ReceiptRequest) (*pb.ReceiptResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.ReceiptResponse), args.Error(1)
}

// Mock GetUsersBySection implementation
func (m *MockServer) GetUsersBySection(ctx context.Context, req *pb.UsersBySectionRequest) (*pb.UsersBySectionResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.UsersBySectionResponse), args.Error(1)
}

// Mock RemoveUser implementation
func (m *MockServer) RemoveUser(ctx context.Context, req *pb.RemoveUserRequest) (*pb.RemoveUserResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.RemoveUserResponse), args.Error(1)
}

// Mock ModifySeat implementation
func (m *MockServer) ModifySeat(ctx context.Context, req *pb.ModifySeatRequest) (*pb.ModifySeatResponse, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*pb.ModifySeatResponse), args.Error(1)
}

// bufconnDialer creates an in-memory connection
func bufconnDialer(listener *bufconn.Listener) func(context.Context, string) (net.Conn, error) {
	return func(ctx context.Context, s string) (net.Conn, error) {
		return listener.Dial()
	}
}

// setupTestServer initializes a mock gRPC server
func setupTestServer(t *testing.T) (*grpc.ClientConn, *MockServer, func()) {
	bufferSize := 1024 * 1024
	listener := bufconn.Listen(bufferSize)

	mockServer := new(MockServer)
	grpcServer := grpc.NewServer()
	pb.RegisterTrainTicketServiceServer(grpcServer, mockServer)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatalf("Failed to start test gRPC server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(context.Background(), "bufnet", grpc.WithContextDialer(bufconnDialer(listener)), grpc.WithInsecure())
	if err != nil {
		t.Fatalf("Failed to create client connection: %v", err)
	}

	cleanup := func() {
		grpcServer.Stop()
		listener.Close()
	}

	return conn, mockServer, cleanup
}

// TestPurchaseTicket verifies ticket purchase
func TestPurchaseTicket(t *testing.T) {
	conn, mockServer, cleanup := setupTestServer(t)
	defer cleanup()

	client := &TrainTicketClient{conn: conn, client: pb.NewTrainTicketServiceClient(conn)}

	mockServer.On("PurchaseTicket", mock.Anything, mock.Anything).Return(&pb.PurchaseResponse{
		Message: "Ticket successfully purchased!",
		Ticket:  &pb.Ticket{From: "CityA", To: "CityB", Price: 100.0, Seat: "A1", Section: "A"},
	}, nil)

	client.PurchaseTicket("CityA", "CityB", "John", "Doe", "john@example.com", 100.0)
}

// TestGetReceipt verifies receipt retrieval
func TestGetReceipt(t *testing.T) {
	conn, mockServer, cleanup := setupTestServer(t)
	defer cleanup()

	client := &TrainTicketClient{conn: conn, client: pb.NewTrainTicketServiceClient(conn)}

	mockServer.On("GetReceipt", mock.Anything, mock.Anything).Return(&pb.ReceiptResponse{
		Message: "Receipt retrieved successfully",
		Receipt: &pb.Receipt{From: "CityA", To: "CityB", PricePaid: 100.0},
	}, nil)

	client.GetReceipt("john@example.com")
}

// TestGetUsersBySection verifies fetching users by section
func TestGetUsersBySection(t *testing.T) {
	conn, mockServer, cleanup := setupTestServer(t)
	defer cleanup()

	client := &TrainTicketClient{conn: conn, client: pb.NewTrainTicketServiceClient(conn)}

	mockServer.On("GetUsersBySection", mock.Anything, mock.Anything).Return(&pb.UsersBySectionResponse{
		Message: "Users in section A retrieved successfully",
		UserSeatInfo: []*pb.UserSeatInfo{
			{User: &pb.User{FirstName: "John", LastName: "Doe", Email: "john@example.com"}, Seat: "A1"},
		},
	}, nil)

	client.GetUsersBySection("A")
}

// TestRemoveUser verifies user removal
func TestRemoveUser(t *testing.T) {
	conn, mockServer, cleanup := setupTestServer(t)
	defer cleanup()

	client := &TrainTicketClient{conn: conn, client: pb.NewTrainTicketServiceClient(conn)}

	mockServer.On("RemoveUser", mock.Anything, mock.Anything).Return(&pb.RemoveUserResponse{
		Success: true,
		Message: "User successfully removed",
	}, nil)

	client.RemoveUser("john@example.com")
}

// TestModifySeat verifies seat modification
func TestModifySeat(t *testing.T) {
	conn, mockServer, cleanup := setupTestServer(t)
	defer cleanup()

	client := &TrainTicketClient{conn: conn, client: pb.NewTrainTicketServiceClient(conn)}

	mockServer.On("ModifySeat", mock.Anything, mock.Anything).Return(&pb.ModifySeatResponse{
		Success:    true,
		NewSeat:    "B3",
		NewSection: "B",
		Message:    "Seat successfully modified",
	}, nil)

	client.ModifySeat("john@example.com", "B3", "B")
}
