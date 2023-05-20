package main

import (
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"main/drivers"
	_ "main/drivers"
	"main/util"
	"os"
	"strings"
)

func main() {

	// get the first prompt if passed
	var initialPrompt = ""
	if len(os.Args) > 1 {
		initialPrompt = strings.Join(os.Args, " ")
	}

	var secretKey, exists = os.LookupEnv("OPENAI_API_KEY")
	if !exists {
		panic("Key not found")
	}
	client := openai.NewClient(secretKey)

	var i = 0
	var prompts []drivers.ChatCompletionRequest
	for {
		i++
		fmt.Printf("### interaction: %d\n", i)
		prompt := drivers.ChatCompletionRequest{}
		if len(initialPrompt) > 0 && i == 1 {
			prompt = drivers.ChatCompletionRequest{Prompt: initialPrompt}
		} else {
			fmt.Print("Enter prompt: ")
			prompt = util.GetInputOrErr()
		}

		prompts = append(prompts, prompt)
		drivers.OpenAIStreamedResponse(client, prompts)

		// for safety, only allow 3 interactions
		if i > 3 {
			break
		}
	}

}
