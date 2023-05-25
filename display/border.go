package display

import gui "github.com/grupawp/warships-gui/v2"

func (f *Fields) DrawBoarder(x, y int) [][]int {
	vec := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, 1},
		{-1, -1},
		{1, -1},
	}
	shipFound := f.FindShip(x, y)
	for _, v := range vec {
		for _, s := range shipFound {
			xA := s[0] + v[0]
			yA := s[1] + v[1]
			if !isInRange(xA, yA) {
				continue
			}
			if !isShip(xA, yA, f) {
				f.Mark(xA, yA, gui.Miss)
			}
		}
	}
	return shipFound
}

func (f *Fields) FindShip(x, y int) [][]int {
	vec := [][]int{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
		{1, 1},
		{-1, 1},
		{-1, -1},
		{1, -1},
	}
	shipPlacement := [][]int{
		{x, y},
	}
	for _, v := range vec {
		shipPlacement = append(shipPlacement, findShipRecursive(x, y, v, f)...)
	}
	return shipPlacement
}
func findShipRecursive(x, y int, v []int, f *Fields) [][]int {
	if x+v[0] < 0 || x+v[0] >= 10 || y+v[1] < 0 || y+v[1] >= 10 {
		return [][]int{}
	}
	if isShip(x+v[0], y+v[1], f) {
		coords := [][]int{{x + v[0], y + v[1]}}
		recursiveCoords := findShipRecursive(coords[0][0], coords[0][1], v, f)
		return append(coords, recursiveCoords...)
	}
	return [][]int{}
}
func isInRange(x, y int) bool {
	return x >= 0 && x < 10 && y >= 0 && y < 10
}
