package service

import (
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"time"
	"warships/domain"
)

type GameService struct {
	PlayerBoard   domain.Board
	OpponentBoard domain.Board
	C             Client
}

func (s *GameService) JoinNewGame() []domain.Mark {
	//start the game

	fmt.Println("joined")
	s.C.StartGame()
	//wait for the game to start
	sts, err := s.C.GetStatus()
	if err != nil {
		return nil
	}
	for sts.GameStatus != "game_in_progress" {
		sts, err = s.C.GetStatus()
		time.Sleep(1 * time.Second)
	}
	//game started
	//get Player Board
	playerBoard, err := s.C.GetBoard()
	if err != nil {
		return nil
	}
	fmt.Println(playerBoard.Board)
	//updatePlayerBaord
	s.PlayerBoard.InitialSetup(mapStringsToMarks(playerBoard.Board))

	fmt.Println(s.PlayerBoard.Marks)
	return s.PlayerBoard.Marks
}

func mapStringsToMarks(s []string) []domain.Mark {
	marks := make([]domain.Mark, len(s))
	for _, v := range s {
		marks = append(marks, mapToMark(v, string(gui.Ship)))
	}
	return marks
}
func mapToMark(coord, state string) domain.Mark {
	if len(coord) > 2 {
		return domain.Mark{
			X:     int(coord[0] - 65),
			Y:     9,
			State: state,
		}
	}
	x := int(coord[1] - 49)
	y := int(coord[0] - 65)
	return domain.Mark{
		X:     x,
		Y:     y,
		State: state,
	}
}
