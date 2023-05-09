package game

import (
	gui "github.com/grupawp/warships-gui/v2"
	"warships/warships/http"
)

type App struct {
	Client
	gui.GUI
	gui.Board
}
type Client interface {
	StartGame() error
	GetBoard() (http.GameBoard, error)
}
