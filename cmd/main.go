package main

import (
	"fmt"
	"warships/warships/game"
)

func main() {
	app := game.New()
	//Display menu
	fmt.Println("Welcome to Warships!")
	fmt.Println("1. Play against WpBot")
	fmt.Println("2. Play against another player")
	fmt.Println("3. List of players")
	fmt.Println("4. Exit")
	fmt.Println("5. Place Ships")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		err := app.Play()
		if err != nil {
			fmt.Println(err)
		}
	case 2:
		fmt.Println("Enter your nickname:")
		var nick string
		fmt.Scanln(&nick)
		fmt.Println("Enter game description:")
		var desc string
		fmt.Scanln(&desc)
		fmt.Println("Enter target nickname:")
		var targetNick string
		fmt.Scanln(&targetNick)
		err := app.PlayPvp(nick, desc, targetNick)
		if err != nil {
			fmt.Println(err)
		}
	case 3:
		app.GetListOfPlayers()
	case 5:
		app.PlaceShips()

	}

}
