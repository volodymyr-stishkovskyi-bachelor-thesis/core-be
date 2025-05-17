package service

import (
	"context"
	"fmt"
	"log"

	"github.com/pinecone-io/go-pinecone/v3/pinecone"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/openai"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/vector"
	"google.golang.org/protobuf/types/known/structpb"
)

type ScrapedDTO struct {
	Text   string
	URL    string
	Source string
}

type QueryDTO struct {
	Query string
	TopK  int
}

func ProcessScraped(ctx context.Context, dto ScrapedDTO) error {
	emb, err := openai.GenerateEmbedding(ctx, dto.Text)
	if err != nil {
		return fmt.Errorf("openai embedding failed: %w", err)
	}

	metadataMap := map[string]interface{}{
		"source": dto.Source,
		"url":    dto.URL,
	}
	metadata, err := structpb.NewStruct(metadataMap)

	vec := &pinecone.Vector{
		Id:       dto.URL,
		Values:   &emb,
		Metadata: metadata,
	}

	if err := vector.UpsertVectors([]*pinecone.Vector{vec}); err != nil {
		return fmt.Errorf("pinecone upsert failed: %w", err)
	}

	if err != nil {
		log.Fatalf("Failed to upsert vectors: %v", err)
	}

	log.Printf("Processed scraped data: %s", dto.URL)
	return nil
}

func RunQuery(ctx context.Context, dto QueryDTO) ([]*pinecone.ScoredVector, error) {
	emb, err := openai.GenerateEmbedding(ctx, dto.Query)
	if err != nil {
		return nil, fmt.Errorf("openai embedding failed: %w", err)
	}

	matches, err := vector.QueryVectors(emb, dto.TopK)
	if err != nil {
		return nil, fmt.Errorf("pinecone query failed: %w", err)
	}

	log.Printf("Query returned %d matches for %q", len(matches), dto.Query)
	return matches, nil
}

func IndexScrapeResponse(ctx context.Context, idPrefix string, resp ScrapeResponse) error {
	var vecs []*pinecone.Vector

	for i, c := range resp.Credly {
		text := fmt.Sprintf("Title: %s; Issuer: %s; IssuedDate: %s", c.Title, c.Issuer, c.IssuedDate)
		emb, err := openai.GenerateEmbedding(ctx, text)
		if err != nil {
			return err
		}
		metadataMapCredly := map[string]interface{}{
			"type":       "credly",
			"title":      c.Title,
			"issuer":     c.Issuer,
			"issuedDate": c.IssuedDate,
		}
		metadataCredly, err := structpb.NewStruct(metadataMapCredly)
		vecs = append(vecs, &pinecone.Vector{
			Id:       fmt.Sprintf("%s:credly:%d", idPrefix, i),
			Values:   &emb,
			Metadata: metadataCredly,
		})
	}

	ll := resp.LeetCode
	statsText := fmt.Sprintf(
		"Reputation: %d; Ranking: %d; Easy submissions: %d; Medium: %d; Hard: %d",
		ll.Reputation, ll.Ranking,
		ll.AcSubmissionNum[0].Count,
		ll.AcSubmissionNum[1].Count,
		ll.AcSubmissionNum[2].Count,
	)
	emb, err := openai.GenerateEmbedding(ctx, statsText)
	if err != nil {
		return err
	}

	metadataMapLeetCode := map[string]interface{}{
		"type":       "leetcode",
		"reputation": fmt.Sprint(ll.Reputation),
		"ranking":    fmt.Sprint(ll.Ranking),
	}
	metadataLeetCode, err := structpb.NewStruct(metadataMapLeetCode)

	vecs = append(vecs, &pinecone.Vector{
		Id:       fmt.Sprintf("%s:leetcode", idPrefix),
		Values:   &emb,
		Metadata: metadataLeetCode,
	})

	if err := vector.UpsertVectors(vecs); err != nil {
		return fmt.Errorf("pinecone upsert failed: %w", err)
	}
	log.Printf("Indexed %d vectors for %s", len(vecs), idPrefix)
	return nil
}
