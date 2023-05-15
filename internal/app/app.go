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
type PlayerList struct {
	Nick       string `json:"nick"`
	GameStatus string `json:"game_status"`
}

type Client interface {
	InitGame(string, string, bool) error
	Board() ([]string, error)
	Status() (*Status, error)
	Fire(coord string) (string, error)
	Description() (*Description, error)
	GetPlayerList() ([]PlayerList, error)
}

// App is the main application
type App struct {
	Client     Client
	board      *gui.Board
	oppShots   []string
	desc       *Description
	shotsHit   uint
	shotsTotal uint
	BotGame    bool
	PlayerName string
	TargetName string
}

// New creates a new instance of App
func New(client Client, board *gui.Board) *App {
	return &App{
		Client: client,
		board:  board,
	}
}

// Run runs the application
func (a *App) Run() (string, error) {
	for {
		a.oppShots = []string{}
		a.board = gui.New(gui.NewConfig())
		fmt.Println("Are you ready?")
		var playAgain bool
		fmt.Scanln(&playAgain)
		for playAgain {
			fmt.Println("Starting the game...")
			err := a.Client.InitGame(a.TargetName, a.PlayerName, a.BotGame)

			if err != nil {
				return "", err
			}

			shipsPlacement, err := a.Client.Board()
			if err != nil {
				return "", err
			}
			_, err = a.WaitForStart()
			if err != nil {
				return "", err
			}

			a.desc, err = a.Client.Description()
			printDescription(*a.desc)

			if err = a.board.Import(shipsPlacement); err != nil {
				return "", err
			}

			if err != nil {
				fmt.Println("Error: ", err)
				return "", err
			}

			play, err := a.Play()
			fmt.Println(play)

			fmt.Println("Would you like to play again?")
			fmt.Scanln(&playAgain)
		}
	}

}
