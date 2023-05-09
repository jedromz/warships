package main

import (
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"warships/warships/game"
	"warships/warships/http"
)

func main() {

	app := game.App{
		Client: &http.HttpClient{},
	}
	err := app.StartGame()
	if err != nil {
		fmt.Println(err)
		return
	}
	board, err := app.GetBoard()
	if err != nil {
		fmt.Println(err)
		return
	}
	mapToState("A1")

	ui := gui.NewGUI(true)

	board1 := gui.NewBoard(1, 1, nil)
	board2 := gui.NewBoard(50, 1, nil)
	ui.Draw(board1)
	ui.Draw(board2)

	states := [10][10]gui.State{}
	for _, v := range board.Board {
		x, y := mapToState(v)
		states[x][y] = gui.Ship
	}
	board1.SetStates(states)

	ui.Start(nil)
}
func mapToState(coord string) (int, int) {
	if len(coord) > 2 {
		return int(coord[0] - 65), 9
	}
	yAxis := int(coord[0] - 65)
	xAxis := int(coord[1] - 49)
	return yAxis, xAxis
}
