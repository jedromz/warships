package game

import (
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"time"
	"warships/warships/boards"
)

func (a *App) GetListOfPlayers() {
	lobby, err := a.GetLobby()
	if err != nil {
		return
	}
	fmt.Println(lobby)
}

func (a *App) GameDescription() {
	dsc, err := a.GetDescription()
	fmt.Println(dsc)
	if err != nil {
		return
	}
	a.Boards.GameDescription = dsc
	a.Boards.GUI.Log(dsc.OppDesc)
}

func (a *App) WaitForGameStart() {
	sts, err := a.GetStatus()
	if err != nil {
		return
	}
	for sts.GameStatus != "game_in_progress" {
		sts, err = a.GetStatus()
		time.Sleep(1 * time.Second)
	}
}

func (a *App) GameLoop() {
	shots := 0
	sts, err := a.GetStatus()
	if err != nil {
		return
	}
	for sts.GameStatus != "ended" {
		txt1 := gui.NewText(60, 1, "Game Status: "+sts.GameStatus, nil)
		a.Boards.GUI.Draw(txt1)
		for _, v := range sts.OppShots[shots:] {
			a.PlayerChannel <- boards.GameEvent{
				v,
				"",
			}
			shots++
		}
		if sts.ShouldFire {
			txt1 := gui.NewText(1, 1, "Fire Away!", nil)
			a.Boards.GUI.Draw(txt1)
			sh := <-a.PlayerShotsChannel
			fire, err := a.Fire(sh.Coords)
			if err != nil {
				return
			}
			a.EnemyChannel <- boards.GameEvent{
				Coords: sh.Coords,
				Result: fire,
			}
		} else {
			txt2 := gui.NewText(1, 1, "Enemy turn", nil)
			a.Boards.GUI.Draw(txt2)
		}
		sts, err = a.GetStatus()
		time.Sleep(1 * time.Second)
	}
	txt1 := gui.NewText(60, 2, "Game Result: "+sts.LastGameStatus, nil)
	a.Boards.GUI.Draw(txt1)
}
