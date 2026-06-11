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
				fmt.Sprintf("You are a helpful assistant that writes clear, concise PR descriptions.\n\nHere is the git diff:\n\n%s\n\nWrite a short PR description with:\n- What changed\n- Why it matters\n\nKeep it under 5 bullet points.", diff),
			)),
		},
	})

	if err != nil {
		fmt.Fprintf(os.Stderr, "API error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(message.Content[0].AsText().Text)
}
