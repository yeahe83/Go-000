package hisdata

import (
	"context"
	"log"
	hisdatapb "xServer/internal/hisdata/api/gen/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements a history data service.
type Service struct {
	hisdatapb.UnimplementedHisdataServiceServer
}

// GetHisData returns history records.
func (s *Service) GetHisData(c context.Context, req *hisdatapb.GetHisDataRequest) (*hisdatapb.GetHisDataResponse, error) {
	log.Fatal("GetHisData")
	return nil, status.Error(codes.Unimplemented, "")
}

// GetRealData returns lastest records.
func (s *Service) GetRealData(c context.Context, req *hisdatapb.GetRealRequest) (*hisdatapb.GetRealResponse, error) {
	log.Fatal("GetRealData")
	return nil, status.Error(codes.Unimplemented, "")
}
