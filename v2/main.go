package main

import "warships/v2/display"

func main() {

	err := display.New().Start()
	if err != nil {
		return
	}

}
