package game

import (
	"fmt"
	"time"
	http "warships/v2/network"
)

type GameService interface { // StartGame starts a new game
	GetGame() (*Game, error)
	Play() error
}

type WpGameService struct {
	g *Game
	c http.Client
}

func NewGameService() *WpGameService {
	return &WpGameService{
		c: &http.HttpClient{},
	}
}
func (w *WpGameService) Play() error {
	fmt.Println("Waiting for game to start...")
	w.c.StartGame()
	//wait for the game to start
	for sts, err := w.c.GetStatus(); sts.GameStatus != "game_in_progress" && err == nil; sts, err = w.c.GetStatus() {
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Game started!")
	return nil
}

func (w *WpGameService) startGame() error {
	return w.c.StartGame()
}
func (w *WpGameService) GetGame() (*Game, error) {
	return w.g, nil
}
