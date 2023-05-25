package display

import (
	gui "github.com/grupawp/warships-gui/v2"
)

type Fields struct {
	States     *[10][10]gui.State
	totalShots int
	totalHits  int
}

func NewFields() *Fields {
	return &Fields{
		States: &[10][10]gui.State{},
	}
}

func (f *Fields) GetAccuracy() float64 {
	if f.totalShots == 0 {
		return 0
	}
	return float64(f.totalHits) / float64(f.totalShots)
}

func (f *Fields) SetStates(s [10][10]gui.State) {
	*f.States = s
}

func (f *Fields) Mark(x, y int, s gui.State) {
	f.States[x][y] = s
}
func (f *Fields) Hit(x, y int) {
	if isShip(x, y, f) {
		f.Mark(x, y, gui.Hit)
	} else {
		f.Mark(x, y, gui.Miss)
	}
}
func isShip(x, y int, f *Fields) bool {
	if x < 0 || x >= 10 || y < 0 || y >= 10 {
		return false
	}
	return f.States[x][y] == gui.Ship || f.States[x][y] == gui.Hit
}
