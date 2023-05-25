package display

import gui "github.com/grupawp/warships-gui/v2"

func mapToState(coord string) (int, int) {
	if len(coord) > 2 {
		return int(coord[0] - 65), 9
	}
	x := int(coord[0] - 65)
	y := int(coord[1] - 49)
	return x, y
}
func stateToMap(x, y int) string {
	if y == 9 {
		return string(x+65) + "10"
	}
	return string(x+65) + string(y+49)
}

func countCells2D(grid [10][10]gui.State) int {
	count := 0
	for _, row := range grid {
		for _, cell := range row {
			if cell != "" {
				count++
			}
		}
	}
	return count
}

func countCells1D(arr []string) int {
	count := 0
	for _, cell := range arr {
		if cell != "" {
			count++
		}
	}
	return count
}
