package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type UploadResponse struct {
	ID string `json:"id"`
}

type AnswerResponse struct {
	Answers []string `json:"answers"`
}

type AnswerRequest struct {
	Model     string   `json:"model"`
	Question  string   `json:"question"`
	Documents []string `json:"documents"`
}

func loadDocuments(wordFile string) string {
	// Read the Word document file
	wordBytes, err := ioutil.ReadFile(wordFile)
	if err != nil {
		panic(err)
	}

	// Encode the Word document as base64
	wordBase64 := base64.StdEncoding.EncodeToString(wordBytes)

	// Prepare the request payload
	payload := map[string]interface{}{
		"title":    "Uploaded Document",
		"document": wordBase64,
		"access":   "private",
	}

	// Convert payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Make the API request to upload the Word document
	apiEndpoint := "https://api.openai.com/v1/documents"
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(payloadJSON))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", "Bearer YOUR_OPENAI_API_KEY")
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read and parse the API response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var uploadResponse UploadResponse
	err = json.Unmarshal(respBody, &uploadResponse)
	if err != nil {
		panic(err)
	}

	// Retrieve the document ID
	documentID := uploadResponse.ID
	fmt.Println("Document ID:", documentID)
}

func main() {
	// API endpoint
	apiEndpoint := "https://api.openai.com/v1/answers"

	// OpenAI API key
	apiKey := "YOUR_OPENAI_API_KEY"

	// Uploaded document IDs
	documentIDs := []string{"doc_1", "doc_2", "doc_3"} // Replace with your actual document IDs

	// Question to ask
	question := "What is the answer to my question?"

	// Prepare the request payload
	payload := AnswerRequest{
		Model:     "davinci",
		Question:  question,
		Documents: documentIDs,
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	// Send the API request
	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(payloadJSON))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("Content-Type", "application/json")

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read and parse the API response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var answerResponse AnswerResponse
	err = json.Unmarshal(respBody, &answerResponse)
	if err != nil {
		panic(err)
	}

	// Extract the answer
	answer := answerResponse.Answers[0]
	fmt.Println("Answer:", answer)
}
