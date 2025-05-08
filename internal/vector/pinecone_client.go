package vector

import (
	"context"
	"log"
	"os"

	"github.com/pinecone-io/go-pinecone/v3/pinecone"
)

var idxConn *pinecone.IndexConnection

func Init() {
	apiKey := os.Getenv("PINECONE_API_KEY")
	host := os.Getenv("PINECONE_INDEX_HOST")

	client, err := pinecone.NewClient(pinecone.NewClientParams{
		ApiKey: apiKey,
	})
	if err != nil {
		log.Fatalf("pinecone: NewClient failed: %v", err)
	}

	idxConn, err = client.Index(pinecone.NewIndexConnParams{
		Host: host,
	})
	if err != nil {
		log.Fatalf("pinecone: connect to index failed: %v", err)
	}

	log.Printf("pinecone: connected to index at %s", host)
}

func UpsertVectors(vectors []*pinecone.Vector) error {
	ctx := context.Background()

	count, err := idxConn.UpsertVectors(ctx, vectors)
	if err != nil {
		return err
	}

	log.Printf("pinecone: upserted %d vectors", count)
	return nil
}

func QueryVectors(queryVector []float32, topK int) ([]*pinecone.ScoredVector, error) {
	ctx := context.Background()

	resp, err := idxConn.QueryByVectorValues(ctx, &pinecone.QueryByVectorValuesRequest{
		TopK:            uint32(topK),
		Vector:          queryVector,
		IncludeMetadata: true,
		IncludeValues:   false,
	})
	if err != nil {
		return nil, err
	}

	return resp.Matches, nil
}
