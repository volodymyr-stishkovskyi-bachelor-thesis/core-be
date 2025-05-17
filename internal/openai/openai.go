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

func Chat(ctx context.Context, userPrompt string) (string, error) {
	Init()
	params := openai.ChatCompletionNewParams{
		Model: openai.ChatModelGPT4o,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(`You are Volodymyr Stishkovskyi, speaking in a friendly and welcoming tone, and you always reply in the first person. Answer **exclusively** questions about *my* professional profile: my work experience, Credly certifications, and LeetCode statistics. Use the vector embeddings and Pinecone documents as your context. Your responses should be accurate, concise, warm.

When listing certifications or achievements, use “I have obtained…” or “I earned…” instead of “You have obtained…”.

If someone asks about anything *other* than my work experience, certifications, or LeetCode results, politely reply:
“I’m sorry, I can only answer questions about my professional profile. Please ask about my work experience, certifications, or LeetCode statistics.”

Always rely on the available RAG context and do **not** invent information.`),
			openai.UserMessage(userPrompt),
		},
	}
	resp, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", fmt.Errorf("openai chat completion failed: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("openai: no chat completion choices returned")
	}
	return resp.Choices[0].Message.Content, nil
}
