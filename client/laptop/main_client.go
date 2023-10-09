package client

import (
	pb "simplebank/pb"

	"google.golang.org/grpc"
)

// LaptopClient is a client to call laptop service RPCs
type LaptopClient struct {
	service pb.LaptopServiceClient
}

// NewLaptopClient returns a new laptop client
func NewLaptopClient(cc *grpc.ClientConn) *LaptopClient {
	service := pb.NewLaptopServiceClient(cc)
	return &LaptopClient{service}
}

// NewLaptopClient returns a new laptop client
func NewServerClient(cc *grpc.ClientConn) *LaptopClient {
	service := pb.NewLaptopServiceClient(cc)
	return &LaptopClient{service}
}
