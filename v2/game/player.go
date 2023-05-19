package game

type Player struct {
	Nick        string
	Description string
	Board       *Board
}

func New(nick, desc string) Player {
	return Player{
		Nick:        nick,
		Description: desc,
		Board:       NewBoard(),
	}
}
