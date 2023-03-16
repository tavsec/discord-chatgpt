package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	AuthToken string
}

func (c *Client) Ask(question string) (ChatCompletion, error) {
	response := ChatCompletion{}
	body := ChatQuestion{
		Model: "gpt-3.5-turbo",
		Messages: []ChatMessage{{
			Role:    "system",
			Content: "You will be answering on Discord, so use only Discord formatting.",
		}, {
			Role:    "user",
			Content: question,
		}},
	}
	s, _ := json.Marshal(body)

	client := http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(s))
	if err != nil {
		return response, err
	}
	req.Header.Add("Authorization", "Bearer "+c.AuthToken)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Err is", err)
	}
	defer res.Body.Close()

	resBody, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(resBody, &response)

	if err != nil {
		return response, err
	}
	return response, nil

}

type ChatQuestion struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatCompletion struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}
