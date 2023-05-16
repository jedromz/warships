package game

import (
	gui "github.com/grupawp/warships-gui/v2"
	"time"
	"warships/warships/dto"
)

type WarshipsController struct {
	C     Client
	State State
}

func (w WarshipsController) GetPlayerBoard() ([10][10]gui.State, error) {
	board, err := w.C.GetBoard()
	if err != nil {
		return [10][10]gui.State{}, err
	}
	err = w.State.UpdatePlayerBoard(getStatesFromGameBoard(board))
	if err != nil {
		return [10][10]gui.State{}, err
	}

	return w.State.GetPlayerBoard(), nil
}

func (w WarshipsController) StartNewGame() error {
	err := w.C.StartGame()
	if err != nil {
		return err
	}

	return nil
}

func (w WarshipsController) WaitForGameStart() (dto.GameStatus, error) {
	var sts dto.GameStatus
	var err error
	for sts, err = w.C.GetStatus(); sts.GameStatus != "game_in_progress"; sts, err = w.C.GetStatus() {
		if err != nil {
			return dto.GameStatus{}, err
		}
		time.Sleep(1 * time.Second)
	}
	w.State.UpdateGameStatus(sts)
	return sts, nil
}
func (w WarshipsController) GetGameDescription() (dto.Description, error) {
	dsc, err := w.C.GetDescription()
	if err != nil {
		return dto.Description{}, err
	}
	w.State.UpdateGameDescription(dsc)
	return dsc, nil
}

func getStatesFromGameBoard(board dto.GameBoard) [10][10]gui.State {
	states := [10][10]gui.State{}
	for _, v := range board.Board {
		x, y := mapToState(v)
		states[x][y] = gui.Ship
	}
	return states
}
func mapToState(coord string) (int, int) {
	if len(coord) > 2 {
		return int(coord[0] - 65), 9
	}
	x := int(coord[1] - 49)
	y := int(coord[0] - 65)

	return x, y
}
