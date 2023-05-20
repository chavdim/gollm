package drivers

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
)

const maxTokens = 500
const model = openai.GPT3Dot5Turbo

type ChatCompletionRequest struct {
	Prompt string
}
type ChatCompletionResponse struct {
	Response string
}

func OpenAIStreamedResponse(client *openai.Client, request []ChatCompletionRequest) {
	var messages []openai.ChatCompletionMessage
	for _, v := range request {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: v.Prompt,
		})
	}
	stream, err := client.CreateChatCompletionStream(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:     model,
			MaxTokens: maxTokens,
			Messages:  messages,
			Stream:    true,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return
	}
	defer stream.Close()

	for {
		var response openai.ChatCompletionStreamResponse
		response, err = stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println()
			return
		}

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return
		}

		fmt.Printf(response.Choices[0].Delta.Content)
	}
}
