package api

import (
	"context"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
)

func initChat(r *gin.RouterGroup) {
	r.POST("/sendChat", chat)
	r.GET("/hello", helloworld)
}

type ChatData struct {
	ApiKey  string
	Content string
}

func chat(ctx *gin.Context) {
	var cd ChatData
	err := ctx.ShouldBind(&cd)
	if err != nil {
		fail(ctx, err)
		return
	}
	if cd.ApiKey == "" {
		//获取本地api_key
	}
	cc := openai.DefaultConfig(cd.ApiKey)
	proxyUrl, _ := url.Parse("http://140.210.195.22:7890")

	cc.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}
	client := openai.NewClientWithConfig(cc)

	rsp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: openai.ChatMessageRoleUser,

				Content: cd.Content,
			},
		},
	})
	if err != nil {
		fail(ctx, err)
		return
	}
	ok(ctx, resp{
		"data": rsp,
	})
}

func helloworld(ctx *gin.Context) {
	ok(ctx, resp{
		"data": "hello world!",
	})
}
