package client

import (
	"context"
	"log"
	"time"

	pb "simplebank/pb"

	"google.golang.org/grpc"
)

// AuthClient is a client to call authentication RPC
type AuthClient struct {
	service   pb.AuthServiceClient
	loginName string
	password  string
}

// NewAuthClient returns a new auth client
func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

// Login login user and returns the access token
func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		LoginName: client.loginName,
		Password:  client.password,
	}

	res, err := client.service.Login(ctx, req)
	log.Printf("Login( %v err: %v", client, err)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}
