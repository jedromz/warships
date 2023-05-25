package display

import (
	gui "github.com/grupawp/warships-gui/v2"
	"testing"
)

func TestCheckPattern(t *testing.T) {
	// Test cases with valid patterns
	validPatterns := []struct {
		coords []string
		result bool
	}{
		{[]string{"A1", "B1", "C1"}, true}, // Horizontal pattern
		{[]string{"A1", "A2", "A3"}, true}, // Vertical pattern
		{[]string{"A1"}, true},
	}

	for _, testCase := range validPatterns {
		result := checkPattern(testCase.coords)
		if result != testCase.result {
			t.Errorf("Expected %v, but got %v for coordinates %v", testCase.result, result, testCase.coords)
		}
	}

	// Test cases with invalid patterns
	invalidPatterns := []struct {
		coords []string
	}{
		{[]string{}},
		{[]string{"A1", "B2", "C4"}},
		{[]string{"A1", "A2", "A3", "B1"}},
		{[]string{"A1", "B1", "C2"}},
	}

	for _, testCase := range invalidPatterns {
		result := checkPattern(testCase.coords)
		if result {
			t.Errorf("Expected false, but got true for coordinates %v", testCase.coords)
		}
	}
}

func TestIsShipPlacementValid_ValidPlacement(t *testing.T) {
	matrix := [10][10]gui.State{
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
		{gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty, gui.Empty},
	}

	// Add ships at valid positions
	matrix[0][0] = gui.Ship
	matrix[0][1] = gui.Ship
	matrix[0][2] = gui.Ship

	if !isShipPlacementValid(matrix) {
		t.Error("Expected ship placement to be valid")
	}
}
