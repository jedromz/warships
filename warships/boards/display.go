package boards

import (
	_ "github.com/JoelOtter/termloop"
	gui "github.com/grupawp/warships-gui/v2"
)

func (b *Boards) DisplayPlayers() {
	playerCnfg := gui.TextConfig{
		FgColor: gui.White,
		BgColor: gui.Green,
	}

	oppCnfg := gui.TextConfig{
		FgColor: gui.White,
		BgColor: gui.Red,
	}
	playerName := gui.NewText(1, 27, b.GameDescription.Nick, &playerCnfg)
	playerDesc := gui.NewText(1, 28, b.GameDescription.Desc, &playerCnfg)

	opponentName := gui.NewText(50, 27, b.GameDescription.Opponent, &oppCnfg)
	opponentDesc := gui.NewText(50, 28, b.GameDescription.OppDesc, &oppCnfg)

	b.GUI.Draw(playerName)
	b.GUI.Draw(playerDesc)
	b.GUI.Draw(opponentName)
	b.GUI.Draw(opponentDesc)

}
func (b *Boards) DrawDashbaord() {

}
