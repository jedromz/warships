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
	urlBoard        = "/api/game/board"
	urlStatus       = "/api/game"
	authTokenHeader = "X-Auth-Token"
	urlFire         = "/api/game/fire"
	urlDesc         = "/api/game/desc"
	urlPlayerList   = "/api/game/list"
)

type FailedToUnmarshalError struct {
	target string
	data   []byte
}

func (e FailedToUnmarshalError) Error() string {
	return fmt.Sprintf("failed to unmarshal %s: %s", e.target, string(e.data))
}

type Board struct {
	Board []string `json:"board"`
}

type Client struct {
	Client *http.Client
	Host   string
	token  string
}

func (c *Client) GetPlayerList() ([]app.PlayerList, error) {
	urlPath, err := url.JoinPath(c.Host, urlPlayerList)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var playerList []app.PlayerList
	err = json.Unmarshal(body, &playerList)
	if err != nil {
		log.Println("failed to unmarshal response", err)
		return nil, err
	}

	return playerList, err
}

func (c *Client) Description() (*app.Description, error) {
	urlPath, err := url.JoinPath(c.Host, urlDesc)

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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response body", err)
		return nil, err
	}

	var desc app.Description
	err = json.Unmarshal(body, &desc)
	if err != nil {
		log.Println("failed to unmarshal response", err)
		return nil, FailedToUnmarshalError{
			target: "app.Description",
			data:   body,
		}
	}

	return &desc, nil
}

func (c *Client) InitGame(targetName string, playerName string, botGame bool) error {
	initGameRequest := map[string]any{
		"coords":      []string{},
		"desc":        "Test Desc",
		"nick":        playerName,
		"target_nick": targetName,
		"wpbot":       botGame,
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failed to read response body", err)
		return nil, err
	}

	fmt.Println(string(body))

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
