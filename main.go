package main

import (
	"log"
	"net"

	"appversal/cron"
	pb "appversal/proto"
	"appversal/server"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reportServer := server.NewReportServer()
	pb.RegisterReportServiceServer(s, reportServer)

	log.Println("gRPC server running on :50051")

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect client: %v", err)
	}
	defer conn.Close()

	client := pb.NewReportServiceClient(conn)

	cron.StartCron(client)

	select {}
}
