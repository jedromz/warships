package boards

import (
	gui "github.com/grupawp/warships-gui/v2"
)

type Fields struct {
	States     *[10][10]gui.State
	TotalShots int
	TotalHits  int
}

func NewFields() *Fields {
	return &Fields{
		States: &[10][10]gui.State{},
	}
}

func (f *Fields) GetAccuracy() float64 {
	if f.TotalShots == 0 {
		return 0
	}
	return float64(f.TotalHits) / float64(f.TotalShots)
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
func (f *Fields) HasShipAround(x, y int) bool {
	// Check all directions for ships
	return isShip(x-1, y, f) || // Left
		isShip(x+1, y, f) || // Right
		isShip(x, y-1, f) || // Up
		isShip(x, y+1, f) || // Down
		isShip(x-1, y-1, f) || // Top Left
		isShip(x-1, y+1, f) || // Bottom Left
		isShip(x+1, y-1, f) || // Top Right
		isShip(x+1, y+1, f) // Bottom Right
}
