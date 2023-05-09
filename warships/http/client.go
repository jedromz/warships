package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GameBoard struct {
	Board []string `json:"board"`
}

type HttpClient struct {
	token string
}

func (c *HttpClient) GetBoard() (GameBoard, error) {
	url := "https://go-pjatk-server.fly.dev/api/game/board"

	// Create the HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return GameBoard{}, err
	}

	// Set the request headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Auth-Token", c.token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return GameBoard{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	var response GameBoard
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error:", err)
		return GameBoard{}, err
	}

	// Extract the board array
	board := response

	// Print the board
	fmt.Println("Board:", board)
	return board, nil
}

func (c *HttpClient) StartGame() error {
	url := "https://go-pjatk-server.fly.dev/api/game"

	body := map[string]interface{}{
		"coords":      []string{},
		"desc":        "My first game",
		"nick":        "John_Doe",
		"target_nick": "",
		"wpbot":       true,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println("Error json:", err)
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error req:", err)
		return err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error resp:", err)
		return err
	}
	defer resp.Body.Close()

	authToken := resp.Header.Get("x-auth-token")

	// Print the auth token
	fmt.Println("Auth Token:", authToken)
	c.token = authToken
	return nil
}
