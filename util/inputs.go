package util

import (
	"bufio"
	"github.com/chavdim/gollm/drivers"
	"os"
)

func GetInputOrErr() drivers.ChatCompletionRequest {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	return drivers.ChatCompletionRequest{Prompt: input}
}
