package main

import (
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
	"net/http"
	"strconv"
	"strings"
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

	a := setUp()

	if ok, _ := strconv.ParseBool(playWithBot); !ok {

		list, _ := a.Client.GetPlayerList()
		printOpponents(list)

		fmt.Println("Choose Your Opponent")
		var oppNumber int

		fmt.Scanln(&oppNumber)

	} else {
		a.TargetName = ""
	}

	a.BotGame, err = strconv.ParseBool(playWithBot)
	a.PlayerName = strings.TrimSpace(playerName)

	run, err := a.Run()
	if err != nil {
		fmt.Println(err)
	}
	printResult(run)
}
func printOpponents(oppList []app.PlayerList) {
	fmt.Println("### LIST OF PLAYERS ###")
	for i, o := range oppList {
		fmt.Println(i, o.Nick)
	}
}
func choosePlayer(oppList []app.PlayerList, n int) string {
	return oppList[n].Nick
}
