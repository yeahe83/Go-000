package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

// GrpcConfig defines a grpc server.
type GrpcConfig struct {
	Name              string
	Addr              string
	AuthPublicKeyFile string
	RegisterFunc      func(*grpc.Server)
}

// RunGrpcServer runs a grpc server.
func RunGrpcServer(c *GrpcConfig) error {

	lis, err := net.Listen("tcp", c.Addr)
	if err != nil {
		log.Fatalf("cannot listen: %v", err)
	}

	// TODO: use grpc.UnaryInterceptor to check token and get user id
	s := grpc.NewServer()
	c.RegisterFunc(s)

	log.Printf("server started %s", c.Addr)
	return s.Serve(lis)
}
