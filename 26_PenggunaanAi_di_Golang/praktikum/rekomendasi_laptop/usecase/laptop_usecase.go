// usecase/laptop_usecase.go
package usecase

import (
	"context"
	openai "github.com/sashabaranov/go-openai"
)

type LaptopUsecase interface {
	RecommendLaptop(userInput, brand, ram, cpu, screenSize, openAIKey string) (string, error)
}

type laptopUsecase struct{}

func NewLaptopUsecase() LaptopUsecase {
	return &laptopUsecase{}
}

func (uc *laptopUsecase) RecommendLaptop(userInput, brand, ram, cpu, screenSize, openAIKey string) (string, error) {
	ctx := context.Background()
	client := openai.NewClient(openAIKey)
	model := openai.GPT3Dot5Turbo
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "Halo, perkenalkan saya sistem untuk rekomendasi laptop",
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: userInput,
		},
	}

	resp, err := uc.getCompletionFromMessages(ctx, client, messages, model)
	if err != nil {
		return "", err
	}
	answer := resp.Choices[0].Message.Content
	return answer, nil
}

func (uc *laptopUsecase) getCompletionFromMessages(
	ctx context.Context,
	client *openai.Client,
	messages []openai.ChatCompletionMessage,
	model string,
) (openai.ChatCompletionResponse, error) {
	if model == "" {
		model = openai.GPT3Dot5Turbo
	}

	resp, err := client.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model:    model,
			Messages: messages,
		},
	)
	return resp, err
}
