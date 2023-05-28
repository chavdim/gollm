[![Go Report Card](https://goreportcard.com/badge/github.com/chavdim/gollm)](https://goreportcard.com/report/github.com/chavdim/gollm)

# GoLLM
Cli wrapper for LLM apis. (currently only supports openAIs chat API) 


![t-rec_1](https://github.com/chavdim/gollm/assets/19467395/bd3032ad-396c-4df2-988d-e665a214ee88)

# Notes
- Expects OpenAIs api key to be set as an environment variable `OPENAI_API_KEY`. see https://platform.openai.com/account/api-keys
- Model, token limits and interactions limits are currently hardcoded. respectively 1000, 100
  - Interactions are currently limited for safety. (loops and billed apis are a scary combination)

# Installation
## homebrew
```
brew install chavdim/taps/gollm   
```
## go
```
go get github.com/chavdim/gollm
```
# Setup
- If no config file exists gollm will ask you to run the config.
- Choose the model to be used and a persona if you want to choose one. 
- Configs are stored in a file ~/.gollm_config.json.

# Usage
## start chat
>gollm
## redo config
>gollm -clean
## start chat with question
>gollm how to assign a variable in golang
> 
>gollm "hello can you tell when will AI overtake humans in intelligence?"
