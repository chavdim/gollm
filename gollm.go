package main

import (
	"flag"
	"github.com/chavdim/gollm/client"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

func main() {
	var forceConfigInitialization = getCleanConfigFlag()
	var config = client.SetupClient(forceConfigInitialization)
	var initialPrompt = getInitialPrompt(forceConfigInitialization)
	var secretKey = attainApiKey()
	var chatClient = client.ChatClient{}
	var openAiClient = openai.NewClient(secretKey)
	chatClient.InitClient(openAiClient, config)
	chatClient.StartChatLoop(initialPrompt)
}

func getCleanConfigFlag() bool {
	// Check if the clean flag is set
	cleanFlag := flag.Bool("clean", false, "run config initialization even if config file exists")
	flag.Parse()
	var forceConfigInitialization = false
	if *cleanFlag {
		forceConfigInitialization = true
	}
	return forceConfigInitialization
}

func getInitialPrompt(forceConfigInitialization bool) string {
	var initialPrompt = ""
	if len(os.Args) > 1 && forceConfigInitialization == false {
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
