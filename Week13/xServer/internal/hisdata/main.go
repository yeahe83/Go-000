package main

import (
	"log"
	hisdatapb "xServer/internal/hisdata/api/gen/v1"
	"xServer/internal/hisdata/hisdata"
	"xServer/internal/shared/server"

	"google.golang.org/grpc"
)

func main() {

	//c := context.Background()

	log.Fatal(server.RunGrpcServer(&server.GrpcConfig{
		Name:              "hisdata",
		Addr:              ":8082",
		AuthPublicKeyFile: "shared/auth/public.key",
		RegisterFunc: func(s *grpc.Server) {
			hisdatapb.RegisterHisdataServiceServer(s, &hisdata.Service{})
		},
	}))
}
