package display

import (
	"github.com/google/uuid"
	"github.com/grupawp/termloop"
)

type flag struct {
}

func newFlag() *flag {
	return &flag{}
}

func (f flag) ID() uuid.UUID {
	return uuid.New()
}

func (f flag) Drawables() []termloop.Drawable {
	//generate random drawables
	d := make([]termloop.Drawable, 0)
	for i := 0; i < 10; i++ {
		d = append(d, termloop.NewRectangle(i*10, i*10, 10, 10, termloop.ColorMagenta))
	}
	return d
}
