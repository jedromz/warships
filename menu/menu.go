package menu

type MenuOption struct {
	Id    int
	Name  string
	Logic func()
}

type Menu struct {
	options []MenuOption
}

func New() Menu {
	return Menu{}
}
func PlayWithBot() MenuOption {
	return MenuOption{
		Id:   1,
		Name: "Play With Bot",
	}
}
