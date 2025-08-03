package chatgpt

import (
	"context"
	"fmt"
	"github.com/assimon/ai-anti-bot/adapter"
	"github.com/assimon/ai-anti-bot/config"
	"github.com/assimon/ai-anti-bot/pkg/json"
	"github.com/assimon/ai-anti-bot/pkg/logger"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"time"
)

var _ adapter.IModel = (*ChatGpt)(nil)

type ChatGpt struct {
	adapter.Option
	Client *openai.Client
}

func NewChatGpt(option adapter.Option) *ChatGpt {
	cfg := openai.DefaultConfig(option.ApiKey)
	if option.Proxy != "" {
		cfg.BaseURL = option.Proxy
	}
	cfg.HTTPClient = &http.Client{
		Timeout: time.Duration(config.Cfg.Ai.Timeout) * time.Second,
	}
	return &ChatGpt{
		Option: option,
		Client: openai.NewClientWithConfig(cfg),
	}
}

func (c *ChatGpt) RecognizeTextMessage(ctx context.Context, userInfo, message string) (adapter.RecognizeResult, error) {
	var result adapter.RecognizeResult
	prompt := fmt.Sprintf(config.Cfg.Prompt.Text, userInfo, message)
	req := openai.ChatCompletionRequest{
		Model: c.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: prompt,
			},
		},
	}
	var resp openai.ChatCompletionResponse
	var err error
	for i := 0; i < config.Cfg.Retry.Times; i++ {
		logger.Log.Infof("Recognize message attempt %d", i+1)
		resp, err = c.Client.CreateChatCompletion(
			ctx,
			req,
		)
		if err == nil {
			logger.Log.Info("Recognize message success")
			break
		}
		logger.Log.Warnf("Recognize message failed: %s", err.Error())
		time.Sleep(time.Duration(config.Cfg.Retry.Delay) * time.Second)
	}
	if err != nil {
		logger.Log.Error("Recognize message finally failed")
		return result, err
	}

	responseContent := resp.Choices[0].Message.Content
	// 兼容处理返回结果是markdown格式的问题
	if len(responseContent) > 7 && responseContent[:7] == "```json" {
		responseContent = responseContent[7 : len(responseContent)-3]
	}
	err = json.C.UnmarshalFromString(responseContent, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (c *ChatGpt) RecognizeImageMessage(ctx context.Context, userInfo, file string) (adapter.RecognizeResult, error) {
	var result adapter.RecognizeResult
	prompt := fmt.Sprintf(config.Cfg.Prompt.Image, userInfo)
	req := openai.ChatCompletionRequest{
		Model: c.Model,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,
				MultiContent: []openai.ChatMessagePart{
					{
						Type: openai.ChatMessagePartTypeText,
						Text: prompt,
					},
					{
						Type: openai.ChatMessagePartTypeImageURL,
						ImageURL: &openai.ChatMessageImageURL{
							URL:    file,
							Detail: openai.ImageURLDetailLow,
						},
					},
				},
			},
		},
	}
	var resp openai.ChatCompletionResponse
	var err error
	for i := 0; i < config.Cfg.Retry.Times; i++ {
		logger.Log.Infof("Recognize message attempt %d", i+1)
		resp, err = c.Client.CreateChatCompletion(
			ctx,
			req,
		)
		if err == nil {
			logger.Log.Info("Recognize message success")
			break
		}
		logger.Log.Warnf("Recognize message failed: %s", err.Error())
		time.Sleep(time.Duration(config.Cfg.Retry.Delay) * time.Second)
	}
	if err != nil {
		logger.Log.Error("Recognize message finally failed")
		return result, err
	}
	responseContent := resp.Choices[0].Message.Content
	// 兼容处理返回结果是markdown格式的问题
	if len(responseContent) > 7 && responseContent[:7] == "```json" {
		responseContent = responseContent[7 : len(responseContent)-3]
	}
	err = json.C.UnmarshalFromString(responseContent, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
