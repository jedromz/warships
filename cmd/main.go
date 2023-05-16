package main

import (
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"warships/warships/game"
	"warships/warships/http"
)

func main() {

	app := game.AppV2{
		C: game.WarshipsController{
			C:     &http.HttpClient{},
			State: &game.WarshipsState{},
		},

		UI: &game.WarshipsUI{
			GUI:           gui.NewGUI(true),
			PlayerBoard:   gui.NewBoard(1, 4, nil),
			OpponentBoard: gui.NewBoard(50, 4, nil),
		},
	}
	err := app.Play()
	if err != nil {
		fmt.Println(err)
	}
}
