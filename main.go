package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(".env error loading: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("API_KEY environment variable must be set")
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	genaiText := genai.Text("What is the meaning of life?")

	iter := model.GenerateContentStream(ctx, genaiText)
	for {
		resp, err := iter.Next()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(resp.Candidates[0].Content)
	}

	// resp, err := model.GenerateContent(ctx, genaiText)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(resp.Candidates[0].Content)
}
