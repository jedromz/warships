package display

import "warships/v2/game"

type WarshipsGui struct {
	gs *game.WpGameService
}

func New() *WarshipsGui {
	return &WarshipsGui{
		gs: game.NewGameService(),
	}
}

func (w *WarshipsGui) Start() error {
	err := w.gs.Play()
	if err != nil {
		return err
	}
	return nil
}
