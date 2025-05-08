package openai

import (
	"context"
	"fmt"
	"os"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

var client *openai.Client

func Init() {
	if client != nil {
		return
	}
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		panic("OPENAI_API_KEY environment variable is not set")
	}

	c := openai.NewClient(option.WithAPIKey(apiKey))
	client = &c
}

func GenerateEmbedding(ctx context.Context, text string) ([]float32, error) {
	Init()

	params := openai.EmbeddingNewParams{
		Input: openai.EmbeddingNewParamsInputUnion{
			OfString: openai.String(text),
		},
		Model: openai.EmbeddingModelTextEmbeddingAda002,
	}

	resp, err := client.Embeddings.New(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("openai: embedding request failed: %w", err)
	}
	if len(resp.Data) == 0 {
		return nil, fmt.Errorf("openai: no embeddings returned")
	}

	raw := resp.Data[0].Embedding
	result := make([]float32, len(raw))
	for i, v := range raw {
		result[i] = float32(v)
	}
	return result, nil
}
