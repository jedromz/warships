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
type Client interface {
	InitGame() error
	Board() ([]string, error)
	Status() (*Status, error)
	Fire(coord string) (string, error)
}

// App is the main application
type App struct {
	client   Client
	board    *gui.Board
	oppShots []string
}

// New creates a new instance of App
func New(client Client, board *gui.Board) *App {
	return &App{
		client: client,
		board:  board,
	}
}

// Run runs the application
func (a *App) Run() error {
	fmt.Println("Starting the game...")
	err := a.client.InitGame()

	if err != nil {
		return err
	}
	shipsPlacement, err := a.client.Board()

	if err != nil {
		return err
	}
	status, err := a.WaitForStart()

	if err != nil {
		return err
	}

	if err = a.board.Import(shipsPlacement); err != nil {
		return err
	}
	printBoard(a.board, status)
	for {
		a.Play()
	}
}
