package game

import (
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"time"
	"warships/warships/http"
)

type PlayerDesc struct {
	nick string
	desc string
}

type App struct {
	Client
	Player        [10][10]gui.State
	Opponent      [10][10]gui.State
	playerBoard   *gui.Board
	opponentBoard *gui.Board
	ui            *gui.GUI
	Shots         chan string
}
type Client interface {
	StartGame() error
	GetBoard() (http.GameBoard, error)
	GetDescription() (http.Description, error)
	Fire(coord string) (string, error)
}

func (a *App) board() {
	txt := gui.NewText(1, 1, "Press on any coordinate to log it.", nil)

	a.ui = gui.NewGUI(true)

	a.playerBoard = gui.NewBoard(1, 1, nil)
	a.opponentBoard = gui.NewBoard(50, 1, nil)

	a.ui.Draw(a.playerBoard)
	a.ui.Draw(a.opponentBoard)

	states := [10][10]gui.State{}

	go func() {
		for {
			char := a.opponentBoard.Listen(context.TODO())
			txt.SetText(fmt.Sprintf("Coordinate: %s", char))
			a.ui.Log("Coordinate: %s", char) // logs are displayed after the game exits
			a.Shots <- char
			a.opponentBoard.SetStates(states)
		}
	}()

	a.ui.Start(nil)
}
func (a *App) Play() error {
	//TODO display board
	go func() {
		a.board()
	}()
	go func() {
		fmt.Println("STARTED!!")
		//TODO start the game
		err := a.StartGame()
		if err != nil {
			return
		}

		//TODO wait for the game to start
		desc, err := a.GetDescription()
		if err != nil {
			return
		}
		for desc.GameStatus != "game_in_progress" {
			time.Sleep(1 * time.Second)
			desc, err = a.GetDescription()
		}
		fmt.Println(desc)

		for desc.GameStatus != "ended" {
			if desc.ShouldFire {

			} else {
				//TODO Load Opponent Shots
				select {
				case msg := <-a.Shots:
					fmt.Println(msg)
				}
			}

		}
	}()

	return nil
}
func mapToState(coord string) (int, int) {
	if len(coord) > 2 {
		return int(coord[0] - 65), 9
	}
	yAxis := int(coord[0] - 65)
	xAxis := int(coord[1] - 49)
	return yAxis, xAxis
}
func UpdateBoard() {

}

//TODO check if my turn to shoot
//TODO check results
//Check game result
