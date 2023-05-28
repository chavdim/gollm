package client

import (
	"fmt"
	"github.com/chavdim/gollm/domain"
	"github.com/chavdim/gollm/drivers"
	"github.com/chavdim/gollm/util"
	"github.com/sashabaranov/go-openai"
)

const interactionLimit = 100

type ChatClient struct {
	client          *openai.Client
	config          domain.Config
	requestHistory  []domain.ChatCompletionRequest
	responseHistory []domain.ChatCompletionResponse
}

func (c *ChatClient) InitClient(client *openai.Client, config domain.Config) {
	c.client = client
	c.config = config
}

func (c *ChatClient) StartChatLoop(initialPrompt string) {
	var i = 0
	for {
		i++
		fmt.Printf("### interaction: %d\n", i)
		var prompt domain.ChatCompletionRequest
		if len(initialPrompt) > 0 && i == 1 {
			prompt = domain.ChatCompletionRequest{Prompt: initialPrompt}
		} else {
			fmt.Print("Enter prompt: ")
			prompt = util.GetInputOrErr()
		}

		c.requestHistory = append(c.requestHistory, prompt)
		var response, err = drivers.OpenAIStreamedResponse(c.client, c.config, c.requestHistory, c.responseHistory)
		c.responseHistory = append(c.responseHistory, response)
		fmt.Printf("\n\n")
		// for safety, limit interactions
		if i == interactionLimit || err != nil {
			break
		}
	}
}
