package main

import (
	"log"
	"xServer/internal/analysis/analysis"
	analysispb "xServer/internal/analysis/api/gen/v1"
	"xServer/internal/shared/server"

	"google.golang.org/grpc"
)

func main() {

	//c := context.Background()

	log.Fatal(server.RunGrpcServer(&server.GrpcConfig{
		Name:              "analysis",
		Addr:              ":8083",
		AuthPublicKeyFile: "shared/auth/public.key",
		RegisterFunc: func(s *grpc.Server) {
			analysispb.RegisterAnalysisServiceServer(s, &analysis.Service{})
		},
	}))
}
