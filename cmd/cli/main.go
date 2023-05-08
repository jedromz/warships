package main

import (
	"flag"
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
	"net/http"
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

	nick := flag.String("nick", "", "player nickname")
	targetNick := flag.String("targetNick", "", "target nickname")
	flag.Parse()

	fmt.Println(*nick)
	fmt.Println(*targetNick)

	a := setUp()
	run, err := a.Run()
	if err != nil {
		fmt.Println(err)
	}
	printResult(run)
}
