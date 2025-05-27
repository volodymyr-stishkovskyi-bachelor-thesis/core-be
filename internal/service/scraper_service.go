package service

import (
	"fmt"

	grpcclient "github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/grpc"
)

type ScraperService interface {
	ConnectAndScrapeHandler() (ScrapeResponse, error)
}

type ScrapeResponse struct {
	Credly   []CredlyResponse `json:"credly"`
	LeetCode LeetCodeStats    `json:"leetcode"`
}

type CredlyResponse struct {
	Title      string `json:"title"`
	Issuer     string `json:"issuer"`
	IssuedDate string `json:"issuedDate"`
}

type LeetCodeStats struct {
	Reputation      int             `json:"reputation"`
	Ranking         int             `json:"ranking"`
	AcSubmissionNum []SubmissionNum `json:"acSubmissionNum"`
}

type SubmissionNum struct {
	Difficulty string `json:"difficulty"`
	Count      int    `json:"count"`
}

func ConnectAndScrapeHandler() (ScrapeResponse, error) {
	scraperClient, conn, err := grpcclient.ConnectToScraperService()
	if err != nil {
		return ScrapeResponse{}, fmt.Errorf("Failed to connect to Scraper Service: %w", err)
	}
	defer conn.Close()

	resp, err := grpcclient.Scrape(scraperClient)
	if err != nil {
		return ScrapeResponse{},
			fmt.Errorf("Failed to scrape: %w", err)
	}
	if resp == nil {
		return ScrapeResponse{},
			fmt.Errorf("Empty response from Scraper Service")
	}

	var result ScrapeResponse

	for _, credlyResp := range resp.Credly {
		result.Credly = append(result.Credly, CredlyResponse{
			Title:      credlyResp.Title,
			Issuer:     credlyResp.Issuer,
			IssuedDate: credlyResp.IssuedDate,
		})
	}

	if resp.Leetcode == nil {
		return result, nil
	}

	var submissions []SubmissionNum
	for _, sub := range resp.Leetcode.AcSubmissionNum {
		submissions = append(submissions, SubmissionNum{
			Difficulty: sub.Difficulty,
			Count:      int(sub.Count),
		})
	}

	result.LeetCode = LeetCodeStats{
		Reputation:      int(resp.Leetcode.Reputation),
		Ranking:         int(resp.Leetcode.Ranking),
		AcSubmissionNum: submissions,
	}

	return result, nil
}
