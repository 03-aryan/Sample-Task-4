package cron

import (
	"context"
	"log"
	"time"

	pb "appversal/proto"

	"github.com/robfig/cron/v3"
)

func StartCron(client pb.ReportServiceClient) {
	c := cron.New()
	userIDs := []string{"user1", "user2", "user3"}

	_, err := c.AddFunc("@every 10s", func() {
		for _, user := range userIDs {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			resp, err := client.GenerateReport(ctx, &pb.GenerateReportRequest{UserId: user})
			if err != nil {
				log.Printf("[%s] Error generating report for %s: %v", time.Now().Format(time.RFC3339), user, err)
			} else if resp.Error != "" {
				log.Printf("[%s] Report error for %s: %s", time.Now().Format(time.RFC3339), user, resp.Error)
			} else {
				log.Printf("[%s] Cron generated report for %s: %s", time.Now().Format(time.RFC3339), user, resp.ReportId)
			}
		}
	})
	if err != nil {
		log.Fatalf("Failed to schedule cron: %v", err)
	}

	c.Start()
}
