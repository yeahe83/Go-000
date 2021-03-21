package analysis

import (
	"context"
	"log"
	analysispb "xServer/internal/analysis/api/gen/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements a history data analysis service.
type Service struct {
	analysispb.UnimplementedAnalysisServiceServer
}

// GetYoY returns year on year data.
func (s *Service) GetYoY(c context.Context, req *analysispb.GetYoYRequest) (*analysispb.GetYoYResponse, error) {

	// TODO: call hisdata by grpc
	// hisdata := ...

	log.Fatal("GetYoY")
	return nil, status.Error(codes.Unimplemented, "")
}

// GetMoM returns month on month data.
func (s *Service) GetMoM(c context.Context, req *analysispb.GetMoMRequest) (*analysispb.GetMoMResponse, error) {
	log.Fatal("GetMoM")
	return nil, status.Error(codes.Unimplemented, "")
}
