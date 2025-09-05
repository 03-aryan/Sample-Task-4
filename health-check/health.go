package main

import (
	"context"
	"fmt"
	"log"

	pb "appversal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect client: %v", err)
	}
	defer conn.Close()

	client := pb.NewReportServiceClient(conn)

	resp, err := client.HealthCheck(context.Background(), &pb.HealthCheckRequest{})
	if err != nil {
		log.Fatalf("HealthCheck error: %v", err)
	}

	fmt.Println(" HealthCheck Response:", resp.Status)
}
