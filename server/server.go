package server

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	pb "appversal/proto"
)


type ReportServer struct {
	pb.UnimplementedReportServiceServer
	mu      sync.Mutex
	reports map[string]string
}

func NewReportServer() *ReportServer {
	return &ReportServer{reports: make(map[string]string)}
}

func (s *ReportServer) GenerateReport(ctx context.Context, req *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error) {
	if req.UserId == "" {
		return &pb.GenerateReportResponse{
			ReportId: "",
			Error:    "UserId cannot be empty",
		}, nil
	}

	reportID := fmt.Sprintf("report-%s-%d", req.UserId, time.Now().UnixNano())

	s.mu.Lock()
	s.reports[reportID] = req.UserId
	s.mu.Unlock()

	log.Printf("[%s] Report generated for UserID %s with ReportID %s",
		time.Now().Format(time.RFC3339), req.UserId, reportID)

	return &pb.GenerateReportResponse{
		ReportId: reportID,
		Error:    "",
	}, nil
}

func (s *ReportServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Status: "Healthy"}, nil
}
