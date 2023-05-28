package game

import (
	"fmt"
	"github.com/google/uuid"
	tl "github.com/grupawp/termloop"
	gui "github.com/grupawp/warships-gui/v2"
)

type RectangleWrapper struct {
	rect *tl.Rectangle
}

type Manual struct {
	id              uuid.UUID
	r               []*RectangleWrapper
	gameDescription gui.Text
}

func (d Manual) Drawables() []tl.Drawable {
	var drawables []tl.Drawable
	for _, r := range d.r {
		drawables = append(drawables, r.rect)
	}
	return drawables
}

func (d Manual) ID() uuid.UUID {
	return d.id
}
func NewDashboardWithRandomRectangles() *Manual {
	dashboard := &Manual{
		id:              uuid.New(),
		r:               []*RectangleWrapper{},
		gameDescription: *gui.NewText(10, 10, string(gui.Ship), nil),
	}
	extraRect := tl.NewRectangle(10, 10, 5, 5, tl.ColorRed)
	extraRectWrapper := &RectangleWrapper{rect: extraRect}
	fmt.Println(dashboard.r)
	dashboard.r = append(dashboard.r, extraRectWrapper)
	fmt.Println(dashboard.r)
	return dashboard
}

func (g *Game) manual() func() {
	return func() {
		fmt.Println(`Battleships Manual
------------------
Welcome to the game of Battleships!
The objective of the game is to sink all the enemy ships.
You will be playing against the computer.
The game board consists of a 10x10 grid.
Each player will place their ships on the grid, and then take turns trying to hit the enemy's ships.
The ships can be placed vertically or horizontally, but not diagonally.
The ships cannot overlap or extend beyond the boundaries of the grid.
You will take turns entering coordinates to target a specific location on the grid.
If you hit an enemy ship, it will be marked as 'S'.
If you miss, it will be marked as 'M'.
The first player to sink all the enemy ships wins!` + gui.Ship + "Ship" + gui.Miss + "Miss" + gui.Hit + "Hit" + gui.Empty + "Empty")
	}
}
