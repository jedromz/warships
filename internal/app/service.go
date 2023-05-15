package app

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	gui "github.com/grupawp/warships-lightgui/v2"
	"os"
	"strings"
	"time"
)

const (
	GameInProgress = "game_in_progress"
	GameOver       = "ended"
	maxRetries     = 10000
	waitTime       = 1 * time.Second
)

type Timer struct {
	playerMove bool
	c          chan string
}

func (t *Timer) display() {

	/*for i := 60; i >= 0; i-- {
		fmt.Printf("Time remaining: %ds\r", i)
		time.Sleep(time.Second)
	}*/

	fmt.Println("Timer finished!")
}

func (a *App) Play() (string, error) {
	state, err := a.Client.Status()
	turn := make(chan string)
	go startTimer(turn)
	for {
		state, err = a.Client.Status()

		if err = a.LoadOppShots(state); err != nil {
			return "", err
		}
		printBoard(a.board, *a.desc)

		fmt.Println((float64(a.shotsHit)/float64(a.shotsTotal))*100, "%")
		if err != nil {
			return "", err
		}

		if state.GameStatus == GameOver {
			break
		}

		if state.ShouldFire {
			a.Fire()
			turn <- "player"
		} else {
			time.Sleep(1 * time.Second)
		}

	}

	return state.LastGameStatus, nil
}

func (a *App) LoadOppShots(state *Status) error {
	a.oppShots = state.OppShots[len(a.oppShots):]
	for _, shot := range a.oppShots {
		mark, err := a.board.HitOrMiss(gui.Left, shot)
		if err != nil {
			return err
		}
		err = a.board.Set(gui.Left, shot, mark)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Fire() {
	shot, err := enterCords()
	if shot == "stop" {
		os.Exit(0)
	}

	fire, err := a.Client.Fire(shot)
	if err != nil {
		return
	}

	switch fire {
	case "miss":
		a.board.Set(gui.Right, shot, gui.Miss)
	case "hit", "sunk":
		a.board.Set(gui.Right, shot, gui.Hit)
		a.shotsHit++
		if fire == "sunk" {
			a.board.CreateBorder(gui.Right, shot)
		}
	}
	a.shotsTotal++
}

func (a *App) WaitForStart() (*Status, error) {
	for retryCount := 0; retryCount < maxRetries; retryCount++ {
		state, err := a.Client.Status()
		if err != nil {
			return nil, err
		}

		if state.GameStatus == GameInProgress {
			return state, nil
		}

		time.Sleep(waitTime)
	}

	return nil, errors.New("unable to start game")
}

func printBoard(board *gui.Board, desc Description) {

	board.Display()
	printDescription(desc)
}
func printDescription(d Description) {
	green := color.New(color.FgHiWhite, color.BgGreen).SprintFunc()
	red := color.New(color.FgHiWhite, color.BgRed).SprintFunc()

	fmt.Printf("%s:%s\n%s:%s\n", green(d.Nick), green(d.Desc), red(d.Opponent), red(d.OppDesc))
}

func enterCords() (string, error) {
	shot := ""
	fmt.Println()
	for !isValidWarshipCoord(shot) {
		fmt.Print("Enter shot: ")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		shot = strings.TrimSpace(text)
	}
	return shot, nil
}

func isValidWarshipCoord(coord string) bool {
	if len(coord) != 2 && len(coord) != 3 {
		return false
	}

	letter := coord[0]
	if letter < 'A' || letter > 'J' {
		return false
	}

	if len(coord) == 3 {
		lastTwoChars := coord[1:]
		if lastTwoChars != "10" {
			return false
		}
	}

	return true
}

func startTimer(input chan string) {
	expiration := time.Now().Add(60 * time.Second)
	timer := time.NewTimer(time.Until(expiration))

	for {
		select {
		case <-timer.C:
			// Timer expired
			fmt.Println("\nTime's up!")
			os.Exit(0)
		case <-time.After(1 * time.Second):
			// Print the time left
			remaining := expiration.Sub(time.Now())
			fmt.Printf("\rTime left: %s", remaining.Round(time.Second))
		case <-input:
			// User entered text, reset the timer
			if !timer.Stop() {
				<-timer.C
			}
			expiration = time.Now().Add(60 * time.Second)
			timer.Reset(time.Until(expiration))
		}
	}
}
