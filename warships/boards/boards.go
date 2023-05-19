package boards

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	gui "github.com/grupawp/warships-gui/v2"
	"strings"
	"time"
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
	ResetTimerChannel  chan bool
	PlayerShots        int
	PlayerHits         int
	Timer              int // Timer in seconds
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

	board1 := gui.NewBoard(1, 5, nil)
	board2 := gui.NewBoard(50, 5, nil)
	rt := make(chan bool)

	return &Boards{
		http.Description{},
		board1,
		board2,
		[10][10]gui.State{},
		[10][10]gui.State{},
		ui,
		c,
		e,
		ps, rt,
		0, 0, 60,
	}

}

func (b *Boards) Display() {

	b.GUI.Draw(b.PlayerBoard)
	b.GUI.Draw(b.EnemyBoard)
	//Enemy board
	go func() {
		states := b.EnemyStates
		for {
			//display accuracy
			b.GUI.Draw(gui.NewText(1, 2, fmt.Sprintf("Accuracy: %.2f %%", b.getPlayerAccuracy()), nil))
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
				b.PlayerHits++
			case "sunk":
				states[y][x] = gui.Hit
				b.PlayerHits++
			}
			b.PlayerShots++
			b.EnemyBoard.SetStates(states)
			b.ResetTimerChannel <- true
		}
	}()

	//Player board
	go func() {
		states := b.PlayerStates
		for {
			event := <-b.PlayerChannel
			y, x := mapToState(event.Coords)
			if b.CheckIfHit(x, y) {
				states[y][x] = gui.Hit
			} else {
				states[y][x] = gui.Miss
			}
			b.ResetTimerChannel <- true
			b.PlayerBoard.SetStates(states)
		}
	}()
	go b.startTimer()
	b.GUI.Start(context.TODO(), nil)

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
func (b *Boards) getPlayerAccuracy() float64 {
	if b.PlayerShots == 0 {
		return 0
	}
	return float64(b.PlayerHits) / float64(b.PlayerShots) * 100
}
func (b *Boards) startTimer() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			b.Timer--

			// Stop the timer at zero
			if b.Timer <= 0 {
				return
			}
			b.GUI.Draw(gui.NewText(1, 3, fmt.Sprintf("Timer: %d seconds", b.Timer), nil))

		case <-b.ResetTimerChannel:
			b.Timer = 60
		}
	}
}

func (b *Boards) PlaceShips() {
	ships := map[int]int{
		4: 1,
		3: 2,
		2: 3,
		1: 4,
	}

	// Display the number of ships to be placed for each ship size

	b.GUI.Draw(b.PlayerBoard)
	go func() {
		states := b.PlayerStates
		for {
			char := b.PlayerBoard.Listen(context.TODO())
			y, x := mapToState(char)
			states[y][x] = gui.Ship
			b.PlayerBoard.SetStates(states)
		}
	}()
	go func() {
		for size, count := range ships {
			b.GUI.Draw(gui.NewText(0, size, fmt.Sprintf("%s x%d\n", strings.Repeat("S", size), count), &gui.TextConfig{
				FgColor: gui.White,
				BgColor: gui.Green,
			}))
		}
	}()
	b.GUI.Start(context.TODO(), nil)
}
