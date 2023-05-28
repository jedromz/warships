package service

import (
	"battleships/boards"
	"battleships/client"
	globals "battleships/globals"
	gui "github.com/grupawp/warships-gui/v2"
	"time"
)

type GameService struct {
	p *boards.Fields
	o *boards.Fields
	c *client.HttpClient
}

func NewGameService() *GameService {
	return &GameService{
		p: boards.NewFields(),
		o: boards.NewFields(),
		c: client.NewHttpClient(),
	}
}

func (s *GameService) GetStats() (globals.GameStats, error) {
	return s.c.GetGameStats()
}

func (s *GameService) GetPlayerAccuracy() float64 {
	return s.o.GetAccuracy()
}
func (s *GameService) GetOpponentAccuracy() float64 {
	return s.p.GetAccuracy()
}

func (s *GameService) MarkOpponent(x, y int, state gui.State) {
	s.o.Mark(x, y, state)
}
func (s *GameService) GetPlayerFields() *boards.Fields {
	return s.p
}
func (s *GameService) GetPlayerStates() *[10][10]gui.State {
	return s.p.States
}
func (s *GameService) GetOpponentFields() *[10][10]gui.State {
	return s.o.States
}

func (s *GameService) StartBotGame(ships []string) error {
	err := s.c.StartGame(ships)
	if err != nil {
		return err
	}
	return nil
}
func (s *GameService) WaitForGame() error {
	for sts, err := s.c.GetStatus(); sts.GameStatus != "game_in_progress" && err == nil; sts, err = s.c.GetStatus() {
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (s *GameService) UpdateGameStatus() (globals.GameStatusResponse, error) {
	sts, err := s.c.GetStatus()
	if err != nil {
		return globals.GameStatusResponse{}, err
	}
	return *sts, nil
}

func (s *GameService) GetDescription() (globals.Description, error) {
	return s.c.GetDescription()
}

func (s *GameService) LoadPlayerBoard() ([10][10]gui.State, error) {
	board, err := s.c.GetBoard()
	if err != nil {
		return [10][10]gui.State{}, err
	}
	var states [10][10]gui.State
	for _, v := range board.Board {
		x, y := mapToState(v)
		states[x][y] = gui.Ship
	}
	s.p.SetStates(states)
	return states, nil
}

func (s *GameService) ListPlayers() ([]globals.LobbyEntry, error) {
	return s.c.GetLobby()
}

func (s *GameService) StartPvpGame(nick, desc, targetNick string, ships []string) {
	s.c.StartPvpGame(nick, desc, targetNick, ships)
}

func (s *GameService) UpdatePlayerStates(states [10][10]gui.State) {
	s.p.SetStates(states)
}

func (s *GameService) UpdatePlayerBoard(coords []string) {
	for _, v := range coords {
		x, y := mapToState(v)
		s.p.Hit(x, y)
	}
}

func (s *GameService) AbortGame() {
	s.c.AbortGame()
}

func (s *GameService) CheckSurroundings(coords []string) bool {

	for _, v := range coords {
		x, y := mapToState(v)
		if s.p.HasShipAround(x, y) {
			return false
		}
	}

	return true
}
