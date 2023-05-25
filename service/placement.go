package service

import (
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
)

func (s *GameService) PlaceShip(coords []string) error {
	fmt.Println(coords)
	if !CheckPattern(coords) {
		fmt.Println("invalid pattern")
		return fmt.Errorf("invalid pattern")
	}
	s.UpdatePlayerBoard(coords)
	return nil
}

func CheckPattern(coords []string) bool {
	if len(coords) == 0 || len(coords) > 4 {
		return false
	}
	x, y := mapToState(coords[0])

	isHorizontal := true
	isVertical := true

	for i := 1; i < len(coords); i++ {
		xi, yi := mapToState(coords[i])

		if xi != x {
			isHorizontal = false
		}
		if yi != y {
			isVertical = false
		}
	}

	if isHorizontal || isVertical {
		return true
	}

	for i := 2; i < len(coords); i++ {
		xi, yi := mapToState(coords[i])
		dx := xi - x
		dy := yi - y

		if (dx == 1 && dy == 0) || (dx == 0 && dy == 1) {

			for j := i + 1; j < len(coords); j++ {
				xj, yj := mapToState(coords[j])
				dx2 := xj - xi
				dy2 := yj - yi

				if (dx2 == dx && dy2 == 0) || (dx2 == 0 && dy2 == dy) {
					return true
				}
			}
		}
	}

	return false
}

var directions = [][]int{
	{0, -1},  // Up
	{0, 1},   // Down
	{-1, 0},  // Left
	{1, 0},   // Right
	{-1, -1}, // Up-left
	{-1, 1},  // Up-right
	{1, -1},  // Down-left
	{1, 1},   // Down-right
}

func IsShipPlacementValid(matrix [10][10]gui.State) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if matrix[i][j] == gui.Ship {
				if !isValidAdjacent(matrix, i, j) {
					return false
				}
			}
		}
	}
	return true
}

func isValidAdjacent(matrix [10][10]gui.State, row, col int) bool {
	for _, dir := range directions {
		adjRow := row + dir[0]
		adjCol := col + dir[1]
		if isValidPosition(adjRow, adjCol) && matrix[adjRow][adjCol] == gui.Ship {
			return false
		}
	}
	return true
}

func isValidPosition(row, col int) bool {
	return row >= 0 && row < 10 && col >= 0 && col < 10
}
