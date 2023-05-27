package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/sashabaranov/go-openai"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename as first argument")

		os.Exit(1)
	}
	filename := os.Args[1]

	gitDiff, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	prompt := `Given the following git diff, which contains the lines changed
and filenames, please provide an appropriate commit message suggestion. Make
sure to highlight any breaking changes explicitly. The commit message should
consist of a subject and a body, separated by two newlines. The subject,
    written in the imperative mood (e.g., "Add", "Fix", "Change"), should be
brief, 50 characters or less. The body of the message should be wrapped at 72
characters.`

	content := fmt.Sprintf("%s\n\n%s", prompt, gitDiff)

	apiKey := os.Getenv("OPENAI_API_KEY")
	if len(apiKey) == 0 {
		fmt.Println("Environment key 'OPENAI_API_KEY' not set.")
		os.Exit(1)
	}

	client := openai.NewClient(apiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       openai.GPT3Dot5Turbo,
			Temperature: 0.2,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("open file %q: %w", filename, err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	sb := strings.Builder{}

	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		if strings.HasPrefix(currentLine, "#") {
			continue
		}
		sb.WriteString(currentLine)
	}

	return sb.String(), nil
}
