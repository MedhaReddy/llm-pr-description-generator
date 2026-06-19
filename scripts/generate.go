package main

import (
	"context"
	"fmt"
	"os"

	anthropic "github.com/anthropics/anthropic-sdk-go"
)

func main() {

	diff := os.Getenv("GET_DIFF")
	if diff == "" {
		diff = `
			+func add(a, b int) int {
			+    return a + b
			+}
		`
	}

	client := anthropic.NewClient()

	message, err := client.Messages.New(context.TODO(), anthropic.MessageNewParams{
		Model:     anthropic.ModelClaudeHaiku4_5,
    	MaxTokens: 1024,
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(
				fmt.Sprintf(`You are a helpful assistant that writes clear, concise PR descriptions.
				Here is the git diff:
				%s
				Write a short PR description with:
				- What changed
				- Why it matters
				Keep it under 5 bullet points.
				Do not use markdown headers (no # symbols). Use only bullet points.`, diff),
			)),
		},
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "API error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(message.Content[0].AsText().Text)
}
