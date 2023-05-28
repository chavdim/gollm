package util

import (
	"bufio"
	"github.com/chavdim/gollm/domain"
	"os"
)

func GetInputOrErr() domain.ChatCompletionRequest {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return domain.ChatCompletionRequest{Prompt: input}
}
