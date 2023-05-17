package domain

type Board struct {
	Marks  []Mark
	Player Player
}

type Mark struct {
	X, Y  int
	State string
}

func (b *Board) Mark(mark Mark) {
	b.Marks = append(b.Marks, mark)
}

func (b *Board) InitialSetup(marks []Mark) {
	b.Marks = marks
}
