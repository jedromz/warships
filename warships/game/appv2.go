package game

import (
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"warships/warships/dto"
)

type AppV2 struct {
	UI UI
	C  Controller
}

type Client interface {
	StartGame() error
	GetBoard() (dto.GameBoard, error)
	GetDescription() (dto.Description, error)
	GetStatus() (dto.GameStatus, error)
	Fire(coord string) (string, error)
}

type State interface {
	GetGameDescription() dto.Description
	GetGameStatus() dto.GameStatus
	GetPlayerBoard() [10][10]gui.State

	UpdateGameStatus(dto.GameStatus)
	UpdateGameDescription(description dto.Description)
	UpdatePlayerBoard([10][10]gui.State) error
	UpdateOpponentBoard([10][10]gui.State) error
}
type UI interface {
	Display([10][10]gui.State)
}
type Controller interface {
	GetPlayerBoard() ([10][10]gui.State, error)
	StartNewGame() error
	WaitForGameStart() (dto.GameStatus, error)
	GetGameDescription() (dto.Description, error)
}

func (a *AppV2) Play() error {
	err := a.C.StartNewGame()
	if err != nil {
		return err
	}
	sts, err := a.C.WaitForGameStart()
	if err != nil {
		return err
	}
	playerBoard, err := a.C.GetPlayerBoard()
	if err != nil {
		return err
	}
	fmt.Println(playerBoard)
	a.UI.Display(playerBoard)
	fmt.Println(sts)
	return nil
}
