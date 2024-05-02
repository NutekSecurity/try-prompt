package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/vertexai/genai"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

var projectId = "shadow-thief"
var region = "us-central1"

func tryGemini(w io.Writer, projectId string, region string, modelName string) error {

	ctx := context.Background()

	credsFile := os.Getenv("GOOGLE_VERTEXAI_CREDENTIALS")
	jsonCreds, err := os.ReadFile(credsFile)
	if err != nil {
		return fmt.Errorf("Failed to read JSON credentials file: %v", err)
	}

	creds, err := google.CredentialsFromJSON(context.Background(), jsonCreds, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return fmt.Errorf("Failed to create credentials from JSON: %v", err)
	}

	var clientOption option.ClientOption
	clientOption = option.WithCredentials(creds)

	client, err := genai.NewClient(ctx, projectId, region, clientOption)

	if err != nil {
		return fmt.Errorf("error creating new client: %w", err)
	}
	model := client.GenerativeModel("gemini-pro")

	resp, err := model.GenerateContent(ctx, genai.Text("What is the average size of a swallow?"))
	if err != nil {
		return fmt.Errorf("error generating content: %w", err)
	}
	rb, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Fprintln(w, string(rb))

	return nil
}

func main() {
	resp := tryGemini(os.Stdout, projectId, region, "projects/shadow-thief/locations/us-central1/models/gemini-1.0-pro")
	if resp != nil {
		fmt.Println(resp)
	} else {
		fmt.Println("Nothing to see here.")
	}
}
