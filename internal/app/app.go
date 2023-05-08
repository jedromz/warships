package app

import (
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
)

// Status is a struct for game status
type Status struct {
	Desc           string   `json:"desc"`
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Nick           string   `json:"nick"`
	OppDesc        string   `json:"opp_desc"`
	OppShots       []string `json:"opp_shots"`
	Opponent       string   `json:"opponent"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}
type Description struct {
	Desc           string   `json:"desc"`
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Nick           string   `json:"nick"`
	OppDesc        string   `json:"opp_desc"`
	OppShots       []string `json:"opp_shots"`
	Opponent       string   `json:"opponent"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}
type Client interface {
	InitGame() error
	Board() ([]string, error)
	Status() (*Status, error)
	Fire(coord string) (string, error)
	Description() (*Description, error)
}

// App is the main application
type App struct {
	client   Client
	board    *gui.Board
	oppShots []string
	desc     *Description
}

// New creates a new instance of App
func New(client Client, board *gui.Board) *App {
	return &App{
		client: client,
		board:  board,
	}
}

// Run runs the application
func (a *App) Run() (string, error) {
	fmt.Println("Starting the game...")
	err := a.client.InitGame()

	if err != nil {
		return "", err
	}

	shipsPlacement, err := a.client.Board()
	if err != nil {
		return "", err
	}
	_, err = a.WaitForStart()
	if err != nil {
		return "", err
	}

	a.desc, err = a.client.Description()
	printDescription(*a.desc)

	if err = a.board.Import(shipsPlacement); err != nil {
		return "", err
	}

	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}

	play, err := a.Play()

	return play, nil
}
