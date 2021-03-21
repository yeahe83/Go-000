package report

import (
	"context"
	"log"
	reportpb "xServer/internal/report/api/gen/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Service implements a report data service.
type Service struct {
	reportpb.UnimplementedReportServiceServer
}

// GetYearReport returns year report.
func (s *Service) GetYearReport(c context.Context, req *reportpb.GetYearReportRequest) (*reportpb.GetYearReportResponse, error) {

	// TODO: call hisdata by grpc
	// hisdata := ...

	log.Fatal("GetYearReportRequest")
	return nil, status.Error(codes.Unimplemented, "")
}

// GetRealData returns month report.
func (s *Service) GetMonthReport(c context.Context, req *reportpb.GetMonthReportRequest) (*reportpb.GetMonthReportResponse, error) {
	log.Fatal("GetMonthReportRequest")
	return nil, status.Error(codes.Unimplemented, "")
}
