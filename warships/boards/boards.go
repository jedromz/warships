package boards

import (
	"context"
	"github.com/google/uuid"
	gui "github.com/grupawp/warships-gui/v2"
	"warships/warships/http"
)

type Boards struct {
	GameDescription    http.Description
	PlayerBoard        *gui.Board
	EnemyBoard         *gui.Board
	PlayerStates       [10][10]gui.State
	EnemyStates        [10][10]gui.State
	GUI                *gui.GUI
	PlayerChannel      chan GameEvent
	EnemyChannel       chan GameEvent
	PlayerShotsChannel chan GameEvent
}
type GameEvent struct {
	Coords string
	Result string
}
type UpdateInfo struct {
	BoardId uuid.UUID
}

func New(c, e, ps chan GameEvent) *Boards {
	ui := gui.NewGUI(true)

	board1 := gui.NewBoard(1, 4, nil)
	board2 := gui.NewBoard(50, 4, nil)

	return &Boards{
		http.Description{},
		board1,
		board2,
		[10][10]gui.State{},
		[10][10]gui.State{},
		ui,
		c,
		e,
		ps,
	}

}

func (b *Boards) Display() {

	b.GUI.Draw(b.PlayerBoard)
	b.GUI.Draw(b.EnemyBoard)

	go func() {
		states := b.EnemyStates
		for {
			char := b.EnemyBoard.Listen(context.TODO())
			b.PlayerShotsChannel <- GameEvent{
				Coords: char,
			}
			event := <-b.EnemyChannel
			y, x := mapToState(event.Coords)
			switch event.Result {
			case "miss":
				states[y][x] = gui.Miss
			case "hit":
				states[y][x] = gui.Hit
			case "sunk":
				states[y][x] = gui.Hit
			}

			b.EnemyBoard.SetStates(states)
		}
	}()

	go func() {
		states := b.PlayerStates
		for {
			event := <-b.PlayerChannel
			y, x := mapToState(event.Coords)
			b.GUI.Log("MAPPED:", y, x)
			b.GUI.Log(event.Coords)
			if b.CheckIfHit(x, y) {
				states[y][x] = gui.Hit
			} else {
				states[y][x] = gui.Miss
			}
			b.PlayerBoard.SetStates(states)
		}
	}()

	b.GUI.Start(nil)

}
func mapToState(coord string) (int, int) {
	if len(coord) > 2 {
		return int(coord[0] - 65), 9
	}
	yAxis := int(coord[0] - 65)
	xAxis := int(coord[1] - 49)
	return yAxis, xAxis
}

func (b *Boards) OpponentBoard(opponentChannel chan string) {

	states := b.EnemyStates
	for {
		char := b.EnemyBoard.Listen(context.TODO())
		y, x := mapToState(char)
		states[y][x] = gui.Hit
		b.EnemyBoard.SetStates(states)
		opponentChannel <- char
	}
}
func (b *Boards) CheckIfHit(x, y int) bool {
	return b.PlayerStates[y][x] == gui.Ship
}
