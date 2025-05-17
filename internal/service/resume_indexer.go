// /internal/service/resume_indexer.go
package service

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/pinecone-io/go-pinecone/v3/pinecone"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/openai"
	"github.com/volodymyr-stishkovskyi-bachelor-thesis/core-be/internal/vector"
	"google.golang.org/protobuf/types/known/structpb"
)

const (
	// размер чанка в символах
	chunkSize = 1000
	// максимально сохраняемая длина текста в метаданных
	maxMetaTextLen = 300
)

// ExtractTextFromPDF открывает PDF и возвращает весь его текст
func ExtractTextFromPDF(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", fmt.Errorf("pdf.Open failed: %w", err)
	}
	defer f.Close()

	reader, err := r.GetPlainText()
	if err != nil {
		return "", fmt.Errorf("GetPlainText failed: %w", err)
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return "", fmt.Errorf("reading plain text failed: %w", err)
	}
	return string(data), nil
}

// chunkText разбивает строку на куски по chunkSize символов, не разрывая слова
func chunkText(text string, size int) []string {
	words := strings.Fields(text)
	var (
		chunks []string
		curr   strings.Builder
	)
	for _, w := range words {
		if curr.Len()+len(w)+1 > size {
			chunks = append(chunks, curr.String())
			curr.Reset()
		}
		if curr.Len() > 0 {
			curr.WriteByte(' ')
		}
		curr.WriteString(w)
	}
	if curr.Len() > 0 {
		chunks = append(chunks, curr.String())
	}
	return chunks
}

// IndexResumePDF извлекает текст из PDF, делит на чанки, эмбеддит и апсертит в Pinecone
func IndexResumePDF(ctx context.Context, idPrefix, pdfPath string) error {
	// 1) Вынимаем текст из PDF
	text, err := ExtractTextFromPDF(pdfPath)
	if err != nil {
		return fmt.Errorf("extract PDF text: %w", err)
	}

	// 2) Разбиваем на чанки
	chunks := chunkText(text, chunkSize)
	if len(chunks) == 0 {
		return fmt.Errorf("no text chunks generated")
	}

	// 3) Подготавливаем векторы
	var vecs []*pinecone.Vector
	for i, chunk := range chunks {
		emb, err := openai.GenerateEmbedding(ctx, chunk)
		if err != nil {
			return fmt.Errorf("embedding chunk %d failed: %w", i, err)
		}

		// создаём превью для метаданных
		preview := chunk
		if len(preview) > maxMetaTextLen {
			preview = preview[:maxMetaTextLen] + "…"
		}

		metadataMap := map[string]interface{}{
			"type":  "resume",
			"chunk": fmt.Sprint(i),
			"text":  preview,
		}
		metadata, err := structpb.NewStruct(metadataMap)

		vecs = append(vecs, &pinecone.Vector{
			Id:       fmt.Sprintf("%s:resume:%d", idPrefix, i),
			Values:   &emb,
			Metadata: metadata,
		})
	}

	// 4) Отправляем все векторы одним Upsert
	if err := vector.UpsertVectors(vecs); err != nil {
		return fmt.Errorf("pinecone upsert failed: %w", err)
	}

	fmt.Printf("Indexed %d resume chunks into Pinecone\n", len(vecs))
	return nil
}
