package main

import (
	"context"
	"log"
	"net/http"

	analysispb "xServer/internal/analysis/api/gen/v1"
	authpb "xServer/internal/auth/api/gen/v1"
	hisdatapb "xServer/internal/hisdata/api/gen/v1"
	reportpb "xServer/internal/report/api/gen/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {

	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()

	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(
			runtime.MIMEWildcard,
			&runtime.JSONPb{
				MarshalOptions: protojson.MarshalOptions{
					UseEnumNumbers: true,
					UseProtoNames:  true,
				},
			},
		),
	)

	serverConfig := []struct {
		name         string
		addr         string
		registerFunc func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)
	}{
		{
			name:         "auth",
			addr:         "localhost:8081",
			registerFunc: authpb.RegisterAuthServiceHandlerFromEndpoint,
		},
		{
			name:         "hisdata",
			addr:         "localhost:8082",
			registerFunc: hisdatapb.RegisterHisdataServiceHandlerFromEndpoint,
		},
		{
			name:         "anaysisdata",
			addr:         "localhost:8083",
			registerFunc: analysispb.RegisterAnalysisServiceHandlerFromEndpoint,
		},
		{
			name:         "reportdata",
			addr:         "localhost:8084",
			registerFunc: reportpb.RegisterReportServiceHandlerFromEndpoint,
		},
	}

	for _, s := range serverConfig {
		err := s.registerFunc(
			c, mux, s.addr,
			[]grpc.DialOption{grpc.WithInsecure()},
		)
		if err != nil {
			log.Fatalf("cannot register service %s: %v", s.name, err)
		}
	}
	addr := ":8080"
	log.Printf("grpc gateway started at %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
