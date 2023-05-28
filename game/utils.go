package game

import (
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"strconv"
	"strings"
	"time"
)

func mapToState(coord string) (int, int) {
	if len(coord) > 2 {
		return int(coord[0] - 65), 9
	}
	x := int(coord[0] - 65)
	y := int(coord[1] - 49)
	return x, y
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
func showSpinner(stop chan bool) {
	spinnerChars := []rune{'|', '-', '/'}

	go func() {
		for i := 0; ; i++ {
			// Clear the console or move the cursor to the beginning of the line (if supported)
			fmt.Print("\033[2K\r")

			// Print the spinner character
			fmt.Printf("Waiting for the game to start %c", spinnerChars[i%len(spinnerChars)])

			// Sleep for a short duration to control the spinner speed
			time.Sleep(100 * time.Millisecond)

			// Check if the stop signal is received
			select {
			case <-stop:
				// Stop the spinner and return
				fmt.Print("\033[2K\r")
				return
			default:
				// Continue spinning
			}
		}
	}()
}

func ValidateBoardPlacement(coords []string) bool {
	if len(coords) < 2 {
		return false
	}

	initialRow, initialCol := int(coords[0][0]-'A'), ParseCoordinate(coords[0][1:])
	deltaRow, deltaCol := int(coords[1][0]-'A')-initialRow, ParseCoordinate(coords[1][1:])-initialCol

	// Diagonal placement is invalid.
	if deltaRow != 0 && deltaCol != 0 {
		return false
	}

	for i := 1; i < len(coords); i++ {
		row := int(coords[i][0] - 'A')
		col := ParseCoordinate(coords[i][1:])

		if row-initialRow != deltaRow*i || col-initialCol != deltaCol*i {
			return false
		}
	}
	return true
}

func ParseCoordinate(coord string) int {
	parsed, err := strconv.Atoi(coord)
	if err != nil {
		panic(err)
	}
	return parsed
}

func isDiagonal(coord1, coord2 string) bool {
	row1, col1 := parseCoord(coord1)
	row2, col2 := parseCoord(coord2)
	return abs(row1-row2) == 1 && abs(col1-col2) == 1
}

func isAdjacent(coord1, coord2 string) bool {
	row1, col1 := parseCoord(coord1)
	row2, col2 := parseCoord(coord2)
	return abs(row1-row2) <= 1 && abs(col1-col2) <= 1
}

func isEdgeTouching(coord string) bool {
	row, col := parseCoord(coord)
	return row == 0 || row == 7 || col == 0 || col == 7
}

func parseCoord(coord string) (int, int) {
	coord = strings.ToUpper(coord)
	row := int(coord[1] - '1')
	col := int(coord[0] - 'A')
	return row, col
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
