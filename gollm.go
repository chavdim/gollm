package main

import (
	"github.com/chavdim/gollm/client"
	"os"
	"strings"
)

func main() {
	var secretKey = attainApiKey()
	var initialPrompt = getInitialPrompt()
	var chatClient = client.ChatClient{}
	chatClient.InitClient(secretKey)
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
