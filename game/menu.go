package game

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type Menu struct {
	options []gameOption
}

func (m Menu) display() {
	for {
		clearScreen()
		fmt.Println("Welcome to Battle Ships!")
		fmt.Println("-------------------------")
		fmt.Println("Please select an option:")
		for i, opt := range m.options {
			fmt.Println(i+1, opt.name)
		}
		fmt.Println("-------------------------")

		var input int
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println(err)
		}

		if input < 0 || input >= len(m.options)+1 {
			fmt.Println("Wrong input")
			return
		}

		fmt.Println("-------------------------")
		m.options[input-1].opt()
		fmt.Println("-------------------------")

		fmt.Println("Press any key to go back to the menu")
		fmt.Scanln()
	}
}

type gameOption struct {
	name string
	opt  func()
}

func (g *Game) Menu() {
	options := []gameOption{
		{name: "Display manual", opt: g.manual()},
		{name: "StartBotGame", opt: g.StartBotGame},
		{name: "StartPvpGame", opt: g.StartPvpGame},
		{name: "ListPlayers", opt: g.listPlayers},
		{name: "Ranking", opt: g.ranking},
	}
	menu := Menu{options: options}
	menu.display()
}

func clearScreen() {
	cmd := exec.Command("clear") // for Linux and macOS
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls") // for Windows
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
