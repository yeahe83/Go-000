package main

import (
	"log"
	reportpb "xServer/internal/report/api/gen/v1"
	"xServer/internal/report/report"
	"xServer/internal/shared/server"

	"google.golang.org/grpc"
)

func main() {

	//c := context.Background()

	log.Fatal(server.RunGrpcServer(&server.GrpcConfig{
		Name:              "report",
		Addr:              ":8084",
		AuthPublicKeyFile: "shared/auth/public.key",
		RegisterFunc: func(s *grpc.Server) {
			reportpb.RegisterReportServiceServer(s, &report.Service{})
		},
	}))
}
