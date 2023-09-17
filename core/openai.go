package core

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const openaiApiUrl string = "https://api.openai.com/v1/chat/completions"

var messages []OpenAiMessage

type OpenAiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type createNewChatRequest struct {
	Model    string          `json:"model"`
	Messages []OpenAiMessage `json:"messages"`
}

type createNewChatResponse struct {
	Choices []struct {
		Index   int
		Message struct {
			Role    string
			Content string
		}
		Finish_reason string
	}
}

func CreateNewChat(message string) string {
	var apiKey = os.Getenv("GPTERMINAL_OPENAI_API_KEY")

	if len(messages) == 0 {
		messages = append(messages, OpenAiMessage{"system", "You are a helpful assistant."})
	}

	messages = append(messages, OpenAiMessage{"user", message})

	jsonDataStruct := createNewChatRequest{"gpt-3.5-turbo", messages}
	jsonData, _ := json.Marshal(jsonDataStruct)

	request, error := http.NewRequest("POST", openaiApiUrl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	request.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	chatResponse := createNewChatResponse{}
	err := json.Unmarshal(body, &chatResponse)

	if err != nil {
		panic(err)
	}

	output := chatResponse.Choices[0].Message.Content
	messages = append(messages, OpenAiMessage{"assistant", output})

	return output
}
