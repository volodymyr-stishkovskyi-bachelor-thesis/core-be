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

func ConnectToLeetCodeService() (pb.LeetCodeServiceClient, *grpc.ClientConn, error) {
	leetcodeUrl := os.Getenv("LEETCODE_URL")
	conn, err := grpc.Dial(leetcodeUrl, grpc.WithInsecure())
	fmt.Println("Connecting to LeetCode Service")
	if err != nil {
		log.Fatalf("Failed to connect to LeetCode Service: %v", err)
		return nil, nil, err
	}

	fmt.Println("Connected to LeetCode Service")
	return pb.NewLeetCodeServiceClient(conn), conn, nil
}

func GetLeetCodeStats(client pb.LeetCodeServiceClient) (*pb.LeetCodeStatsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.Empty{}
	resp, err := client.GetStats(ctx, req)
	if err != nil {
		log.Fatalf("Error getting LeetCode stats: %v", err)
		return nil, err
	}
	log.Printf("Retrieved LeetCode stats successfully")

	return resp, nil
}
