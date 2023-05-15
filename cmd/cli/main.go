package main

import (
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
	"net/http"
	"strconv"
	"warships/internal/app"
	"warships/internal/client"
)

func setUp() *app.App {
	c := client.Client{
		Client: &http.Client{},
		Host:   "https://go-pjatk-server.fly.dev",
	}
	board := gui.New(gui.NewConfig())
	a := app.New(
		&c,
		board,
	)
	return a
}
func printResult(run string) {
	switch run {
	case "lose":
		fmt.Println("You lost...")
	case "win":
		fmt.Println("You won!")
	default:
		fmt.Println("Hmmmm..")
	}
}
func main() {

	fmt.Println("Play with bot? y/n: ")
	var playWithBot string
	_, err := fmt.Scanln(&playWithBot)
	if err != nil {
		return
	}
	fmt.Println("Enter your name:")
	var playerName string
	_, err = fmt.Scanln(&playerName)
	if err != nil {
		return
	}

	fmt.Println("Enter target name:")
	var targetName string
	_, err = fmt.Scanln(&targetName)
	/*
		nick := flag.String("nick", "", "player nickname")
		targetNick := flag.String("targetNick", "", "target nickname")
		playWithBot := flag.String("bot", "", "start a game vs bot")
		flag.Parse()

		fmt.Println(*nick)
		fmt.Println(*targetNick)
		fmt.Println(*playWithBot)
	*/

	a := setUp()

	list, err := a.Client.GetPlayerList()
	fmt.Println(list)

	a.BotGame, err = strconv.ParseBool(playWithBot)
	a.PlayerName = playerName
	a.TargetName = targetName
	run, err := a.Run()
	if err != nil {
		fmt.Println(err)
	}
	printResult(run)
}
