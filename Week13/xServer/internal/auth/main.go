package main

import (
	"log"
	authpb "xServer/internal/auth/api/gen/v1"
	"xServer/internal/auth/auth"
	"xServer/internal/shared/server"

	"google.golang.org/grpc"
)

func main() {

	//c := context.Background()

	log.Fatal(server.RunGrpcServer(&server.GrpcConfig{
		Name:              "auth",
		Addr:              ":8081",
		AuthPublicKeyFile: "shared/auth/public.key",
		RegisterFunc: func(s *grpc.Server) {
			authpb.RegisterAuthServiceServer(s, &auth.Service{})
		},
	}))
}
