package drivers

import (
	"context"
	"errors"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
)

const maxTokens = 1000
const model = openai.GPT3Dot5Turbo

type ChatCompletionRequest struct {
	Prompt string
}
type ChatCompletionResponse struct {
	Response string
}

func OpenAIStreamedResponse(client *openai.Client, requestHistory []ChatCompletionRequest, responseHistory []ChatCompletionResponse) (ChatCompletionResponse, error) {
	var messages []openai.ChatCompletionMessage
	totalMessages := len(requestHistory) + len(responseHistory)
	for i := 0; i < totalMessages; i++ {
		var v string
		var message openai.ChatCompletionMessage
		if i%2 == 0 {
			v, requestHistory = requestHistory[0].Prompt, requestHistory[1:]
			message = openai.ChatCompletionMessage{Role: openai.ChatMessageRoleUser, Content: v}
		} else {
			v, responseHistory = responseHistory[0].Response, responseHistory[1:]
			message = openai.ChatCompletionMessage{Role: openai.ChatMessageRoleAssistant, Content: v}
		}
		messages = append(messages, message)
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
		return ChatCompletionResponse{""}, err
	}
	defer stream.Close()
	var chosenResponses []string
	for {
		var response openai.ChatCompletionStreamResponse
		response, err = stream.Recv()

		if errors.Is(err, io.EOF) {
			var result = strings.Join(chosenResponses, "")
			return ChatCompletionResponse{result}, nil
		}
		var chosen = response.Choices[0].Delta.Content
		chosenResponses = append(chosenResponses, chosen)

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return ChatCompletionResponse{""}, err
		}
		fmt.Print(chosen)
	}
}
