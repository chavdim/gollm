# gollm
[![Go Report Card](https://goreportcard.com/badge/github.com/chavdim/gollm)](https://goreportcard.com/report/github.com/chavdim/gollm)

# GoLLM
Cli wrapper for LLM apis. (currently only supports openAI chat API) 

# Notes
- Interactions are currently limited to 5 for safety.
- Expects OpenAIs api key to be set as an environment variable `OPENAI_API_KEY`. see https://platform.openai.com/account/api-keys
- Model, token limits and interactions limits are currently hardcoded. respectively GPT3.5, 500, 5

# Instalation
```
go get github.com/chavdim/gollm
```
# Usage
## start chat
>gollm 
## start chat with question
>gollm how to assign a variable in golang
> 
>gollm "hello can you tell when will AI overtake humans in intelligence?"
