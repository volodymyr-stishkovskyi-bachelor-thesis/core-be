package grpcclient

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
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

func Scrape(client pb.ScraperServiceClient) (*pb.ScrapeResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &emptypb.Empty{}
	resp, err := client.Scrape(ctx, req)
	if err != nil {
		log.Fatalf("Error using Scrape: %v", err)
		return nil, err
	}
	log.Printf("Title: %s", resp.GetCredly()[0])
	log.Printf("Title: %s", resp.GetLeetcode())

	return resp, nil
}
