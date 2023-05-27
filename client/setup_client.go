package client

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/sashabaranov/go-openai"
)

type Config struct {
	model   string
	persona string
}

func SetupClient() {
	prompt := promptui.Select{
		Label: "Select model",
		Items: []string{openai.GPT3Dot5Turbo, openai.GPT4},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	fmt.Printf("You choose %q\n", result)
}
