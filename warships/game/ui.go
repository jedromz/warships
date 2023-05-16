package game

import (
	gui "github.com/grupawp/warships-gui/v2"
)

type WarshipsUI struct {
	GUI           *gui.GUI
	PlayerBoard   *gui.Board
	OpponentBoard *gui.Board
}

func (w *WarshipsUI) Display(playerStates [10][10]gui.State) {

	w.GUI.Draw(w.PlayerBoard)
	w.GUI.Draw(w.OpponentBoard)

	w.PlayerBoard.SetStates(playerStates)

	go func() {
		for {

		}
	}()

	go func() {

	}()

	w.GUI.Start(nil)

}
