package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"warships/internal/app"
)

const (
	applicationJson = "application/json"
	urlInitGame     = "/api/game"
	urlBoard        = "/api//game/board"
	urlStatus       = "/api/game"
	authTokenHeader = "X-Auth-Token"
	urlFire         = "/api/game/fire"
)

type Board struct {
	Board []string `json:"board"`
}

type Client struct {
	Client *http.Client
	Host   string
	token  string
}

func (c *Client) InitGame() error {
	initGameRequest := map[string]any{
		"coords":     nil,
		"desc":       "",
		"nick":       "",
		"targetNick": "",
		"wpbot":      true,
	}

	b, err := json.Marshal(initGameRequest)
	if err != nil {
		return fmt.Errorf("client#InitGame: failed to marshal request: %v", err)
	}
	urlPath, err := url.JoinPath(c.Host, urlInitGame)
	if err != nil {
		return fmt.Errorf("client#InitGame: failed to create request: %v", err)
	}
	resp, err := http.Post(urlPath, applicationJson, bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("client#InitGame: failed to post: %v", err)
	}
	defer resp.Body.Close()
	c.token = resp.Header.Get(authTokenHeader)
	return nil
}

func (c *Client) Board() ([]string, error) {
	urlPath, err := url.JoinPath(c.Host, urlBoard)

	fmt.Println(urlPath)
	if err != nil {
		fmt.Errorf("error")
	}

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Println("failed to create request", err)
	}

	req.Header.Set(authTokenHeader, c.token)

	resp, err := c.Client.Do(req)
	if err != nil {
		log.Println("failed to get", err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	fmt.Println(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response body", err)
		return nil, err
	}

	var board Board
	err = json.Unmarshal(body, &board)
	if err != nil {
		log.Println("failed to unmarshal response", err)
		return nil, err
	}

	return board.Board, nil
}

func (c *Client) Status() (*app.Status, error) {
	urlPath, err := url.JoinPath(c.Host, urlStatus)
	if err != nil {
		fmt.Errorf("error")
	}

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		log.Println("failed to create request", err)
	}

	req.Header.Set(authTokenHeader, c.token)

	resp, err := c.Client.Do(req)
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read body", err)
	}

	var status app.Status
	err = json.Unmarshal(result, &status)
	if err != nil {
		log.Println("failed to unmarshal", err)
	}

	return &status, nil
}

func (c *Client) Fire(coord string) (string, error) {
	cord := struct {
		Coord string `json:"coord"`
	}{
		Coord: coord,
	}
	b, err := json.Marshal(cord)

	urlPath, err := url.JoinPath(c.Host, urlFire)
	if err != nil {
		fmt.Errorf("error")
	}

	req, err := http.NewRequest(http.MethodPost, urlPath, bytes.NewReader(b))
	if err != nil {
		log.Println("failed to create request", err)
	}

	req.Header.Set(authTokenHeader, c.token)

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	result, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var body map[string]string
	err = json.Unmarshal(result, &body)
	if err != nil {
		log.Println("failed to unmarshal", err)
	}
	return body["result"], nil
}