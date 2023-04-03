package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type APIRequest struct {
	Model       string `json:"model"`
	Prompt      string `json:"prompt"`
	MaxTokens   int    `json:"max_tokens"`
	Temperature int    `json:"temperature"`
	Stream      bool   `json:"stream"`
}

type APIResponse struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func CallAPI(question string, config Config, transport *http.Transport) (string, error) {
	client := &http.Client{
		Transport: transport,
	}

	apiRequest := APIRequest{
		Model:       config.Model,
		Prompt:      question,
		MaxTokens:   config.MaxTokens,
		Temperature: config.Temperature,
		Stream:      config.Stream,
	}

	requestBody, err := json.Marshal(apiRequest)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", config.APIURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.APIKey)
	req.Header.Set("Accept", "text/event-stream")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	reader := bufio.NewReader(resp.Body)
	var apiResponse APIResponse
	var answer string

	for {
		// 创建一个通道接收一行数据
		lineChan := make(chan string, 1)

		// 使用 time.AfterFunc 设置超时
		timeout := time.AfterFunc(time.Duration(config.APITimeout)*time.Second, func() {
			lineChan <- ""
		})

		// 在一个新的 Goroutine 中读取一行数据
		go func() {
			line, err := reader.ReadString('\n')
			if err == nil {
				lineChan <- line
			}
		}()

		// 等待从 lineChan 接收到数据
		line := <-lineChan

		// 停止超时处理
		timeout.Stop()

		if line == "" {
			if answer != "" {
				//log.Println("API timed out, partial answer:", answer)
				return answer, nil
			}
			//log.Println("API timed out, no answer")
			return "", errors.New("API timed out")
		}

		if strings.HasPrefix(line, "data:") {
			line = strings.TrimSpace(strings.TrimPrefix(line, "data:"))
			//log.Println("Received data:", line)

			if line == "[DONE]" {
				return answer, nil
			}

			err = json.Unmarshal([]byte(line), &apiResponse)
			if err != nil {
				return "", err
			}

			if len(apiResponse.Choices) > 0 {
				answerPart := strings.TrimSpace(apiResponse.Choices[0].Text)
				if answerPart != "" {
					answer += " " + answerPart
					fmt.Print(answerPart)
				}
			}
		}
	}
}
