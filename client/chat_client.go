package client

import (
	"fmt"
	"github.com/chavdim/gollm/drivers"
	"github.com/chavdim/gollm/util"
	"github.com/sashabaranov/go-openai"
)

const interactionLimit = 5

type ChatClient struct {
	client         *openai.Client
	messageHistory []drivers.ChatCompletionRequest
}

func (c *ChatClient) InitClient(client *openai.Client) {
	c.client = client
}

func (c *ChatClient) StartChatLoop(initialPrompt string) {
	var i = 0
	for {
		i++
		fmt.Printf("### interaction: %d\n", i)
		var prompt drivers.ChatCompletionRequest
		if len(initialPrompt) > 0 && i == 1 {
			prompt = drivers.ChatCompletionRequest{Prompt: initialPrompt}
		} else {
			fmt.Print("Enter prompt: ")
			prompt = util.GetInputOrErr()
		}

		c.messageHistory = append(c.messageHistory, prompt)
		drivers.OpenAIStreamedResponse(c.client, c.messageHistory)

		// for safety, limit interactions
		if i == interactionLimit {
			break
		}
	}
}
