package main

import (
	gui "github.com/grupawp/warships-gui/v2"
	"warships/warships/game"
	"warships/warships/http"
)

func main() {
	c := make(chan string)
	app := game.App{
		Client:   &http.HttpClient{},
		Player:   [10][10]gui.State{},
		Opponent: [10][10]gui.State{},
	}
	app.Play()
	<-c

	///////////////////////////////

}

/*

	app := game.App{
		Client:   &http.HttpClient{},
		Player:   [10][10]gui.State{},
		Opponent: [10][10]gui.State{},
	}

	app.StartGame()
	board, _ := app.GetBoard()
	for _, coord := range board.Board {
		y, x := mapToState(coord)
		app.Player[x][y] = gui.Ship
	}

	txt := gui.NewText(1, 1, "Press on any coordinate to log it.", nil)

	ui := gui.NewGUI(true)

	board1 := gui.NewBoard(1, 1, nil)
	board1.SetStates(app.Player)
	board2 := gui.NewBoard(50, 1, nil)

	ui.Draw(board1)
	ui.Draw(board2)

	states := [10][10]gui.State{}

	go func() {
		for {
			char := board2.Listen(context.TODO())
			txt.SetText(fmt.Sprintf("Coordinate: %s", char))
			ui.Log("Coordinate: %s", char) // logs are displayed after the game exits
			x, y := mapToState(char)
			fire, _ := app.Fire(char)
			switch fire {
			case "hit":
				states[x][y] = gui.Hit
			case "miss":
				states[x][y] = gui.Miss
			case "sunk":
				states[x][y] = gui.Hit

			}
			board2.SetStates(states)
		}
	}()

	ui.Start(nil)
*/
