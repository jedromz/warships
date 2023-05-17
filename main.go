package main

import (
	gui2 "github.com/grupawp/warships-gui/v2"
	"warships/domain"
	"warships/service"
	"warships/shipsgui"
)

func main() {
	srvc := service.GameService{
		PlayerBoard: domain.Board{
			Marks:  nil,
			Player: domain.Player{},
		},
		OpponentBoard: domain.Board{
			Marks:  nil,
			Player: domain.Player{},
		},
		C: &service.HttpClient{},
	}
	gui := shipsgui.GUI{
		Service:       srvc,
		GUI:           gui2.NewGUI(true),
		PlayerBoard:   gui2.NewBoard(1, 4, nil),
		OpponentBoard: gui2.NewBoard(1, 4, nil),
	}

	gui.Display()

}
