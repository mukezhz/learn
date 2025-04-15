package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdaurl"
	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/googlegenai"
)

func main() {
	lambda.Start(lambdaurl.Wrap(http.HandlerFunc(handler)))
}

// ChunkData represents a response chunk that will be sent through the channel
type ChunkData struct {
	Content string
	Error   error
	Done    bool
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Get prompt from query parameter 'q'
	prompt := r.URL.Query().Get("q")
	if prompt == "" {
		http.Error(w, "Prompt cannot be empty", http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	chunkChan := make(chan ChunkData)

	// go DummyStreamResponse(ctx, chunkChan, prompt)

	go StreamFromGemini(ctx, prompt, chunkChan)

	for chunk := range chunkChan {
		if chunk.Error != nil {
			fmt.Printf("Stream error: %v\n", chunk.Error)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if chunk.Done {
			fmt.Println("Stream finished")
			return
		}

		data := fmt.Sprintf("data: %s\n\n", chunk.Content)
		if _, err := w.Write([]byte(data)); err != nil {
			fmt.Printf("Write error: %v\n", err)
			return
		}
		fmt.Printf("Sent chunk: %s\n", chunk.Content)
	}
	w.(http.Flusher).Flush()
}

func StreamFromGemini(ctx context.Context, prompt string, chunkChan chan<- ChunkData) {
	defer close(chunkChan)

	g, err := genkit.Init(ctx,
		genkit.WithPlugins(&googlegenai.GoogleAI{}),
		genkit.WithDefaultModel("googleai/gemini-2.0-flash"),
	)
	if err != nil {
		chunkChan <- ChunkData{Error: err}
		return
	}

	fmt.Printf("Starting to stream response for prompt: %s\n", prompt)

	stream, err := genkit.Generate(ctx, g,
		ai.WithPrompt(prompt),
		ai.WithStreaming(func(ctx context.Context, chunk *ai.ModelResponseChunk) error {
			text := chunk.Text()
			if text != "" {
				chunkChan <- ChunkData{Content: text}
			}
			return nil
		}),
	)
	if err != nil {
		chunkChan <- ChunkData{Error: err}
		return
	}

	chunkChan <- ChunkData{
		Content: stream.Text(),
	}

	chunkChan <- ChunkData{
		Done: true,
	}
}

func DummyStreamResponse(ctx context.Context, chunkChan chan<- ChunkData, prompt string) {
	defer close(chunkChan)

	dummyResponses := []string{
		"Hello there! ",
		"This is a sample ",
		"response from the ",
		"Lambda function. ",
		"It's simulating ",
		"a streaming response ",
		"with multiple chunks. ",
		"Your prompt was: ",
		"Yoo " + prompt,
	}

	for _, text := range dummyResponses {
		chunkChan <- ChunkData{Content: text}
		time.Sleep(200 * time.Millisecond)
	}

	chunkChan <- ChunkData{Done: true}
}
