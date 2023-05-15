package game

import (
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
}

type Client interface {
	StartGame() error
	GetBoard() (http.GameBoard, error)
	GetDescription() (http.Description, error)
	GetStatus() (*http.GameStatusResponse, error)
	Fire(coord string) (string, error)
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

	for {

	}
	return nil
}

func (a *App) SetUpBoard(board http.GameBoard) {
	for _, coords := range board.Board {
		y, x := mapToState(coords)
		a.Boards.PlayerStates[y][x] = gui.Ship
	}
}
func mapToState(coord string) (int, int) {
	if len(coord) > 2 {
		return int(coord[0] - 65), 9
	}
	yAxis := int(coord[0] - 65)
	xAxis := int(coord[1] - 49)
	return yAxis, xAxis
}
