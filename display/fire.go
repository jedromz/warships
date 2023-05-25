package display

import (
	gui "github.com/grupawp/warships-gui/v2"
)

func (s *GameService) Fire(coords string) error {
	res, err := s.c.Fire(coords)
	if err != nil {
		return err
	}
	x, y := mapToState(coords)
	switch res {
	case "hit":
		s.o.totalHits++
		s.o.Mark(x, y, gui.Hit)
	case "miss":
		s.o.Mark(x, y, gui.Miss)
	case "sunk":
		s.o.totalHits++
		s.o.Mark(x, y, gui.Hit)
		s.o.DrawBoarder(x, y)
	}
	s.o.totalShots++
	return nil
}
