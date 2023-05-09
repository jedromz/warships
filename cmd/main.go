package main

import (
	gui "github.com/grupawp/warships-gui/v2"
)

func main() {
	ui := gui.NewGUI(true)

	board := gui.NewBoard(1, 1, nil)
	ui.Draw(board)
	ui.Start(nil)
}
