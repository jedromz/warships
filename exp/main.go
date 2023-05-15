package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func() {
		seconds := 0
		for {
			fmt.Printf("\rTimer: %d seconds", seconds)
			seconds++
			time.Sleep(1 * time.Second)
		}
	}()

	var userInput string
	go func() {
		for {
			fmt.Print("\nEnter something: ")
			fmt.Scanln(&userInput)
			fmt.Println("You entered:", userInput)
		}
	}()

	<-c
}
