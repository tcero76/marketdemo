package config

import (
	"os"

	anthropic "github.com/anthropics/anthropic-sdk-go"
)

type AnthropicClient struct {
	Client anthropic.Client
}

func NewAI() *AnthropicClient {
	model := os.Getenv("ANTHROPIC_MODEL")

	if model == "" {
		model = "claude-sonnet-5"
	}

	return &AnthropicClient{
		Client: anthropic.NewClient(),
	}
}
