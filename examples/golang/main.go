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
	// Check if a file name has been provided as an argument to the program
	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename as first argument")
		os.Exit(1)
	}
	filename := os.Args[1]

	// Read the provided file and retrieve the git diff
	gitDiff, err := readFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Prepare the content for the OpenAI API by formatting the prompt
	content := formatPrompt(gitDiff)

	// Check if OpenAI API key is available in environment variables
	apiKey := os.Getenv("OPENAI_API_KEY")
	if len(apiKey) == 0 {
		fmt.Println("Environment key 'OPENAI_API_KEY' not set.")
		os.Exit(1)
	}

	// Initialize the OpenAI client using the API key
	client := openai.NewClient(apiKey)

	// Request a chat completion from OpenAI API
	responseContent, err := createChatCompletion(client, content)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	// Print the received response content
	fmt.Println(responseContent)
}

// readFile reads the file at the given path and filters out lines beginning with '#'
func readFile(filename string) (string, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("open file %q: %w", filename, err)
	}
	defer file.Close()

	// Prepare a new scanner for the file
	fileScanner := bufio.NewScanner(file)
	sb := strings.Builder{}

	// Loop through each line in the file
	for fileScanner.Scan() {
		currentLine := fileScanner.Text()
		// Add the line to a string builder if it doesn't start with '#'
		if !strings.HasPrefix(currentLine, "#") {
			sb.WriteString(currentLine)
		}
	}

	// Check if there were any errors during scanning
	if err := fileScanner.Err(); err != nil {
		return "", fmt.Errorf("reading file %q: %w", filename, err)
	}

	// Return the accumulated text
	return sb.String(), nil
}

// formatPrompt formats the content to be sent to OpenAI for chat completion
func formatPrompt(gitDiff string) string {
	// Define the prompt
	prompt := `Given the following git diff, which contains the lines changed
and filenames, please provide an appropriate commit message suggestion. Make
sure to highlight any breaking changes explicitly. The commit message should
consist of a subject and a body, separated by two newlines. The subject,
written in the imperative mood (e.g., "Add", "Fix", "Change"), should be
brief, 50 characters or less. The body of the message should be wrapped at 72
characters.`

	// Combine the prompt with the git diff
	return fmt.Sprintf("%s\n\n%s", prompt, gitDiff)
}

// createChatCompletion sends the request to OpenAI for chat completion and returns the response content
func createChatCompletion(client *openai.Client, content string) (string, error) {
	// Make the request for chat completion
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
		return "", err
	}

	// Return the content of the response
	return resp.Choices[0].Message.Content, nil
}
