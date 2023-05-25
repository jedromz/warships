package service

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
		s.o.TotalHits++
		s.o.Mark(x, y, gui.Hit)
	case "miss":
		s.o.Mark(x, y, gui.Miss)
	case "sunk":
		s.o.TotalHits++
		s.o.Mark(x, y, gui.Hit)
		s.o.DrawBoarder(x, y)
	}
	s.o.TotalShots++
	return nil
}
