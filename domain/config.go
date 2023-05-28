package domain

type Config struct {
	Model              string `koanf:"model"`
	PersonaDescription string `koanf:"persona"`
}
