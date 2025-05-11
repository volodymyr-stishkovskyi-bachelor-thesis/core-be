package grpcclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/proto"
	"google.golang.org/grpc"
)

func ConnectToScraperService() (pb.ScraperServiceClient, *grpc.ClientConn, error) {
	scraperUrl := os.Getenv("SCRAPER_URL")
	conn, err := grpc.Dial(scraperUrl, grpc.WithInsecure())
	fmt.Println("Connecting to Scraper Service")
	if err != nil {
		log.Fatalf("Failed to connect to Scraper Service: %v", err)
		return nil, nil, err
	}

	fmt.Println("Connected to Scraper Service")
	return pb.NewScraperServiceClient(conn), conn, nil
}

func Scrape(url string, client pb.ScraperServiceClient) (*pb.ScrapeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.ScrapeRequest{
		Url: url,
	}
	resp, err := client.Scrape(ctx, req)
	if err != nil {
		log.Fatalf("Error using Scrape: %v", err)
		return nil, err
	}
	log.Printf("Title: %s", resp.GetTitle())

	return resp, nil
}
