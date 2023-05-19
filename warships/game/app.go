package game

import (
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"warships/warships/boards"
	"warships/warships/http"
)

type App struct {
	Client
	Boards             *boards.Boards
	PlayerChannel      chan boards.GameEvent
	EnemyChannel       chan boards.GameEvent
	PlayerShotsChannel chan boards.GameEvent
	hits               int
	totalShots         int
}

type Client interface {
	StartGame() error
	StartPvpGame(nick, desc, targetNick string) error
	GetBoard() (http.GameBoard, error)
	GetDescription() (http.Description, error)
	GetStatus() (*http.GameStatusResponse, error)
	Fire(coord string) (string, error)
	GetLobby() ([]http.LobbyEntry, error)
}

func New() App {
	c := make(chan boards.GameEvent)
	e := make(chan boards.GameEvent)
	ps := make(chan boards.GameEvent)
	return App{
		Client:             &http.HttpClient{},
		Boards:             boards.New(c, e, ps),
		PlayerChannel:      c,
		EnemyChannel:       e,
		PlayerShotsChannel: ps,
	}
}

func (a *App) Play() error {
	err := a.StartGame()
	if err != nil {
		return err
	}
	a.WaitForGameStart()
	board, err := a.GetBoard()
	if err != nil {
		return err
	}
	//Wait for the game to start
	a.WaitForGameStart()

	a.GameDescription()
	a.Boards.DisplayPlayers()
	a.SetUpBoard(board)

	a.Boards.PlayerBoard.SetStates(a.Boards.PlayerStates)
	go a.Boards.Display()
	a.GameLoop()

	return nil
}

func (a *App) PlayPvp(nick, desc, targetNick string) error {
	fmt.Println("here1")
	err := a.StartPvpGame(nick, desc, targetNick)
	if err != nil {
		return err
	}
	board, err := a.GetBoard()
	fmt.Println("here2")

	if err != nil {
		return err
	}
	//Wait for the game to start
	a.WaitForGameStart()
	fmt.Println("here3")

	a.GameDescription()
	a.Boards.DisplayPlayers()
	a.SetUpBoard(board)

	a.Boards.PlayerBoard.SetStates(a.Boards.PlayerStates)
	go a.Boards.Display()
	a.GameLoop()

	return nil
}

func (a *App) SetUpBoard(board http.GameBoard) {
	for _, coords := range board.Board {
		y, x := mapToState(coords)
		a.Boards.PlayerStates[y][x] = gui.Ship
	}
}
func (a *App) PlaceShips() {
	a.Boards.PlaceShips()
}
func mapToState(coord string) (int, int) {
	if len(coord) > 2 {
		return int(coord[0] - 65), 9
	}
	yAxis := int(coord[0] - 65)
	xAxis := int(coord[1] - 49)
	return yAxis, xAxis
}
