package shipsgui

import (
	"context"
	gui "github.com/grupawp/warships-gui/v2"
	"warships/service"
)

type GUI struct {
	Service       service.GameService
	GUI           *gui.GUI
	PlayerBoard   *gui.Board
	OpponentBoard *gui.Board
}

func (g *GUI) Display() {
	marks := g.Service.JoinNewGame()

	ui := gui.NewGUI(true)

	board1 := gui.NewBoard(1, 4, nil)
	board2 := gui.NewBoard(50, 4, nil)

	ui.Draw(board1)
	ui.Draw(board2)

	states := [10][10]gui.State{}
	for _, v := range marks {
		states[v.X][v.Y] = gui.State(v.State)
	}
	board1.SetStates(states)
	board2.SetStates(states)

	go g.listenOpponentBoard()
	ui.Start(context.TODO())
}

func (g *GUI) listenOpponentBoard() {
	go func() {
		for {
			char := g.OpponentBoard.Listen(context.TODO())
			g.GUI.Log("Coordinate: %s", char) // logs are displayed after the game exits
		}
	}()
}
