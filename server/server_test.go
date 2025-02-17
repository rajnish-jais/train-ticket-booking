package server

import (
	"net"
	"testing"

	"google.golang.org/grpc"
	pb "train-ticket-booking/proto"

	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":50051")
	assert.NoError(t, err)
	assert.NotNil(t, lis)

	grpcServer := grpc.NewServer()
	assert.NotNil(t, grpcServer)

	pb.RegisterTrainTicketServiceServer(grpcServer, nil)
}
