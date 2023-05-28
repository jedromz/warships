package service

import (
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
)

func (s *GameService) PlaceShip(coords []string) error {
	if !CheckPattern(coords) {
		return fmt.Errorf("invalid pattern")
	}
	if !s.CheckSurroundings(coords) {
		return fmt.Errorf("invalid placement")
	}
	s.UpdatePlayerBoard(coords)
	return nil
}

func CheckPattern(coords []string) bool {
	if len(coords) == 0 || len(coords) > 4 {
		return false
	}

	x, y := mapToState(coords[0])
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} // Right, Left, Down, Up

	for i := 1; i < len(coords); i++ {
		xi, yi := mapToState(coords[i])

		if xi != x && yi != y {
			return false // Invalid coordinate placement (not horizontal or vertical)
		}

		found := false
		for _, dir := range directions {
			dx, dy := dir[0], dir[1]
			if xi == x+dx && yi == y+dy {
				found = true
				break
			}
		}

		if !found {
			return false // Invalid coordinate placement (not adjacent)
		}

		x, y = xi, yi
	}

	return true
}

func isValidPlacement(coords []string, index, x, y int) bool {
	if index >= len(coords) {
		return true
	}

	xi, yi := mapToState(coords[index])
	dx := xi - x
	dy := yi - y

	if isValidLShape(dx, dy) && isValidPosition(xi, yi) {
		// Recursively check the remaining coordinates
		return isValidPlacement(coords, index+1, xi, yi)
	}

	return false
}

func isValidLShape(dx, dy int) bool {
	return (dx == 1 && dy == 0) ||
		(dx == 0 && dy == 1) ||
		(dx == -1 && dy == 0) ||
		(dx == 0 && dy == -1)
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
				if !isValidShipPlacement(matrix, i, j) {
					return false
				}
			}
		}
	}
	return true
}

func isValidShipPlacement(matrix [10][10]gui.State, row, col int) bool {
	// Check if the adjacent cells contain a ship
	for _, dir := range directions {
		adjRow := row + dir[0]
		adjCol := col + dir[1]
		if isValidPosition(adjRow, adjCol) && matrix[adjRow][adjCol] == gui.Ship {
			return false
		}
	}

	// Recursively check the diagonal cells if they contain a ship
	for _, dir := range diagonalDirections {
		adjRow := row + dir[0]
		adjCol := col + dir[1]
		if isValidPosition(adjRow, adjCol) && matrix[adjRow][adjCol] == gui.Ship {
			// Check the orthogonal neighbors of the diagonal cell
			if !isValidOrthogonalPlacement(matrix, adjRow, adjCol) {
				return false
			}
		}
	}

	return true
}

func isValidOrthogonalPlacement(matrix [10][10]gui.State, row, col int) bool {
	// Check if the orthogonal neighbors of the diagonal cell contain a ship
	for _, dir := range orthogonalDirections {
		adjRow := row + dir[0]
		adjCol := col + dir[1]
		if isValidPosition(adjRow, adjCol) && matrix[adjRow][adjCol] == gui.Ship {
			return false
		}
	}

	return true
}

var orthogonalDirections = [][]int{
	{0, -1}, // Left
	{0, 1},  // Right
	{-1, 0}, // Up
	{1, 0},  // Down
}

var diagonalDirections = [][]int{
	{-1, -1}, // Up-left
	{-1, 1},  // Up-right
	{1, -1},  // Down-left
	{1, 1},   // Down-right
}

func isValidPosition(row, col int) bool {
	return row >= 0 && row < 10 && col >= 0 && col < 10
}
