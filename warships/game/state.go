package game

import (
	gui "github.com/grupawp/warships-gui/v2"
	"warships/warships/dto"
)

type WarshipsState struct {
	Status        dto.GameStatus
	Description   dto.Description
	PlayerBoard   [10][10]gui.State
	OpponentBoard [10][10]gui.State
}

func (w *WarshipsState) GetGameStatus() dto.GameStatus {
	return w.Status
}
func (w *WarshipsState) GetGameDescription() dto.Description {
	return w.Description
}
func (w *WarshipsState) GetPlayerBoard() [10][10]gui.State {
	return w.PlayerBoard
}

func (w *WarshipsState) UpdateGameStatus(sts dto.GameStatus) {
	w.Status = sts
}

func (w *WarshipsState) UpdateGameDescription(description dto.Description) {
	w.Description = description
}

func (w *WarshipsState) UpdatePlayerBoard(states [10][10]gui.State) error {
	w.PlayerBoard = states
	return nil
}

func (w *WarshipsState) UpdateOpponentBoard(states [10][10]gui.State) error {
	//TODO implement me
	panic("implement me")
}
