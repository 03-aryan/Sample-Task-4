package main

import (
    "context"
    "log"
    "time"

    pb "appversal/proto"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials/insecure"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewReportServiceClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    resp, err := client.GenerateReport(ctx, &pb.GenerateReportRequest{UserId: "user1"})
    if err != nil {
        log.Fatalf("Error calling GenerateReport: %v", err)
    }
    log.Printf("GenerateReport response: %v", resp)

    healthResp, err := client.HealthCheck(ctx, &pb.HealthCheckRequest{})
    if err != nil {
        log.Fatalf("Error calling HealthCheck: %v", err)
    }
    log.Printf("HealthCheck response: %v", healthResp)
}
