package main

import (
	gui "github.com/grupawp/warships-gui/v2"
)

func main() {

	ui := gui.NewGUI(true)

	board1 := gui.NewBoard(1, 1, nil)
	board2 := gui.NewBoard(50, 1, nil)
	ui.Draw(board1)
	ui.Draw(board2)

	states := [10][10]gui.State{}
	for i := range states {
		states[i] = [10]gui.State{}
	}
	board1.SetStates(states)

	ui.Start(nil)
}

