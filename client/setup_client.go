package client

import (
	"fmt"
	"github.com/chavdim/gollm/domain"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"github.com/manifoldco/promptui"
	"github.com/sashabaranov/go-openai"
	"os"
)

var k = koanf.New(".")
var parser = json.Parser()
var homeDir, _ = os.UserHomeDir()
var configPath = fmt.Sprintf("%s/.gollm_config.json", homeDir)

// SetupClient initializes the client and returns the config.
// if forceInitialize is true, it will run the setup process even if the config file exists.
func SetupClient(forceInitialize bool) domain.Config {
	if forceInitialize {
		selectedConfig := RunSetupConfig()
		WriteConf(selectedConfig)
	}

	if err := k.Load(file.Provider(configPath), json.Parser()); err != nil {
		fmt.Println("Error loading domain:", err)
		selectedConfig := RunSetupConfig()
		WriteConf(selectedConfig)
	}

	var config domain.Config
	_ = k.Unmarshal("", &config)
	return config
}

func RunSetupConfig() domain.Config {
	// setup model
	promptModel := promptui.Select{
		Label: "Select model",
		Items: []string{openai.GPT3Dot5Turbo, openai.GPT4},
	}
	_, selectedModel, err := promptModel.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
	}
	fmt.Printf("You choose %q\n", selectedModel)

	// setup persona
	var personaNames []string
	for _, persona := range domain.Personas {
		personaNames = append(personaNames, persona.Name)
	}
	promptPersona := promptui.Select{
		Label: "Select PersonaDescription",
		Items: personaNames,
	}
	_, selectedPersona, err := promptPersona.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
	}
	fmt.Printf("You choose %q\n", selectedPersona)

	var personaDescription string
	for _, persona := range domain.Personas {
		if persona.Name == selectedPersona {
			personaDescription = persona.Description
		}
	}
	return domain.Config{Model: selectedModel, PersonaDescription: personaDescription}
}

func WriteConf(config domain.Config) {
	_ = k.Set("model", config.Model)
	_ = k.Set("persona", config.PersonaDescription)
	// Open a file to write
	f, err := os.Create(configPath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer f.Close()
	s, _ := k.Marshal(parser)
	// Write the JSON data to the file
	_, err = f.Write(s)
	if err != nil {
		fmt.Println("Error writing JSON data to file:", err)
		return
	}
}
