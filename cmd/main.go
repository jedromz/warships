package main

import (
	"warships/warships/game"
)

func main() {

	app := game.New()

	app.Play()

	//fmt.Println(<-app.PlayerChannel)
}
