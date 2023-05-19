package game

import gui "github.com/grupawp/warships-gui/v2"

type Board struct {
	States [10][10]gui.State
}

func NewBoard() *Board {
	return &Board{
		States: [10][10]gui.State{},
	}
}
func NewBoardFromStates(states [10][10]gui.State) *Board {
	return &Board{
		States: states,
	}
}
