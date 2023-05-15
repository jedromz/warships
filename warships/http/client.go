package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GameBoard struct {
	Board []string `json:"board"`
}
type FireResponse struct {
	Coord string `json:"coord"`
}

type Description struct {
	Nick           string   `json:"nick"`
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Opponent       string   `json:"opponent"`
	OppShots       []string `json:"opp_shots"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}
type HttpClient struct {
	token string
}

func (c *HttpClient) GetDescription() (Description, error) {
	url := "https://go-pjatk-server.fly.dev/api/game"
	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return Description{}, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-Auth-Token", c.token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Description{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Description{}, err
	}

	var desc Description
	err = json.Unmarshal(body, &desc)
	if err != nil {
		return Description{}, err
	}
	return desc, nil
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

func (c *HttpClient) Fire(coord string) (string, error) {
	url := "https://go-pjatk-server.fly.dev/api/game/fire"

	// Create the request body
	requestBody := map[string]string{
		"coord": coord,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Token", c.token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	var response map[string]string
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// Get the "result" value from the response
	result := response["result"]

	// Print the fire result
	return result, nil
}
