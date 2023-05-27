package main

import (
	"github.com/chavdim/gollm/client"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func main() {
	client.SetupClient()
	var secretKey = attainApiKey()
	var initialPrompt = getInitialPrompt()
	var chatClient = client.ChatClient{}
	var openAiClient = openai.NewClient(secretKey)
	chatClient.InitClient(openAiClient)
	chatClient.StartChatLoop(initialPrompt)
}

func getInitialPrompt() string {
	var initialPrompt = ""
	if len(os.Args) > 1 {
		initialPrompt = strings.Join(os.Args, " ")
	}
	return initialPrompt
}

func attainApiKey() string {
	var secretKey, exists = os.LookupEnv("OPENAI_API_KEY")
	if !exists {
		panic("Key not found. Set OPENAI_API_KEY env var see https://platform.openai.com/account/api-keys")
	}
	return secretKey
}
