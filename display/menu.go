package display

import "fmt"

type Menu struct {
	options []gameOption
}

func (m Menu) display() {
	for i, opt := range m.options {
		fmt.Println(i, opt.name)
	}
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}
	if input < 0 || input >= len(m.options) {
		fmt.Println("Wrong input")
		return
	}
	m.options[input].opt()
	return
}

type gameOption struct {
	name string
	opt  func()
}

func (g *Game) Menu() {
	options := []gameOption{
		{name: "StartBotGame", opt: g.StartBotGame},
		{name: "ListPlayers", opt: g.listPlayers},
		{name: "StartPvpGame", opt: g.StartPvpGame},
	}
	menu := Menu{options: options}
	menu.display()
}
