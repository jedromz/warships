package main

import (
	"fmt"
	"time"
)

func main() {
	timer(60)
}

func timer(seconds int) {
	for i := seconds; i >= 0; i-- {
		fmt.Printf("Time remaining: %ds\r", i)
		time.Sleep(time.Second)
	}

	fmt.Println("Timer finished!")
}
