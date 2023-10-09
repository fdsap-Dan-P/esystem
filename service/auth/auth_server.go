package service

import (
	"context"
	"log"

	pb "simplebank/pb"

	dsUsr "simplebank/db/datastore/user"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServiceServer is the server for authentication
type AuthServiceServer struct {
	userStore  dsUsr.QueriesUser //.UserStore
	jwtManager *JWTManager
	pb.UnimplementedAuthServiceServer
}

// NewAuthServiceServer returns a new auth server
func NewAuthServiceServer(userStore dsUsr.QueriesUser, jwtManager *JWTManager) *AuthServiceServer {
	a := pb.UnimplementedAuthServiceServer{}
	return &AuthServiceServer{userStore, jwtManager, a}
}

// Login is a unary RPC to login user
func (server *AuthServiceServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Printf("Login %v", req.GetLoginName())
	user, err := server.userStore.GetUserbyName(ctx, req.GetLoginName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	// log.Println("Login 1", user.LoginName)

	if !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}
	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	res := &pb.LoginResponse{AccessToken: token}
	return res, nil
}
