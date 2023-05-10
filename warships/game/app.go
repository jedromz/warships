package game

import (
	gui "github.com/grupawp/warships-gui/v2"
	"warships/warships/http"
)

type PlayerDesc struct {
	nick string
	desc string
}

type App struct {
	Client
	gui.GUI
	gui.Board
	player   string
	opponent string
}
type Client interface {
	StartGame() error
	GetBoard() (http.GameBoard, error)
	GetDescription() error
}
