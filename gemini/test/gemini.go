package main

import (
	"context"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"log"
)

const (
	API_KEY = "xxxx"
)

func main() {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	//client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("API_KEY")))
	client, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatalf("err1: %v", err)
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	// 1.http
	//resp, err := model.GenerateContent(ctx, genai.Text("使用中文讲一个鬼故事"))
	//if err != nil {
	//	log.Fatalf("err2: %v", err)
	//}
	//log.Printf("PromptFeedback: %+v", resp.PromptFeedback)
	//for _, c := range resp.Candidates {
	//	log.Printf("Candidate: %+v", c.Content)
	//}

	// 2.流式响应
	iter := model.GenerateContentStream(ctx, genai.Text("使用中文讲一个鬼故事"))

	for {
		resp, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("err3: %v", err)
		}
		for _, c := range resp.Candidates {
			for _, item := range c.Content.Parts {
				fmt.Printf("%s\n", item)
			}

		}
	}
}
