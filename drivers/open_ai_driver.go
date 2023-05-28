package drivers

import (
	"context"
	"errors"
	"fmt"
	"github.com/chavdim/gollm/domain"
	"github.com/sashabaranov/go-openai"
	"io"
	"strings"
)

const maxTokens = 1000

func OpenAIStreamedResponse(
	client *openai.Client,
	config domain.Config,
	requestHistory []domain.ChatCompletionRequest,
	responseHistory []domain.ChatCompletionResponse,
) (domain.ChatCompletionResponse, error) {
	var messages []openai.ChatCompletionMessage
	// apply persona if exists
	if len(config.PersonaDescription) > 0 {
		message := openai.ChatCompletionMessage{Role: openai.ChatMessageRoleSystem, Content: config.PersonaDescription}
		messages = append(messages, message)
	}
	// apply all previous prompts as responses
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
			Model:     config.Model,
			MaxTokens: maxTokens,
			Messages:  messages,
			Stream:    true,
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletionStream error: %v\n", err)
		return domain.ChatCompletionResponse{""}, err
	}
	defer stream.Close()
	var chosenResponses []string
	for {
		var response openai.ChatCompletionStreamResponse
		response, err = stream.Recv()

		if errors.Is(err, io.EOF) {
			var result = strings.Join(chosenResponses, "")
			return domain.ChatCompletionResponse{result}, nil
		}
		var chosen = response.Choices[0].Delta.Content
		chosenResponses = append(chosenResponses, chosen)

		if err != nil {
			fmt.Printf("\nStream error: %v\n", err)
			return domain.ChatCompletionResponse{""}, err
		}
		fmt.Print(chosen)
	}
}
