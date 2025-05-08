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
