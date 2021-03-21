package auth

import (
	"context"
	"log"
	authpb "xServer/internal/auth/api/gen/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements a auth service.
type Service struct {
	authpb.UnimplementedAuthServiceServer
}

// Login logs a user.
func (s *Service) Login(c context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	log.Fatal("Login")
	return nil, status.Error(codes.Unimplemented, "")
}
