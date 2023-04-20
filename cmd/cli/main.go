package main

import (
	"fmt"
	gui "github.com/grupawp/warships-lightgui/v2"
	"net/http"
	"warships/internal/app"
	"warships/internal/client"
)

func main() {
	c := client.Client{
		Client: &http.Client{},
		Host:   "https://go-pjatk-server.fly.dev",
	}
	board := gui.New(gui.NewConfig())
	a := app.New(
		&c,
		board,
	)
	err := a.Run()
	if err != nil {
		fmt.Println(err)
	}
}
