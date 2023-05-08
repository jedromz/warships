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
	maxRetries     = 5
	waitTime       = 1 * time.Second
)

func (a *App) Play() (string, error) {
	state, err := a.client.Status()
	for {
		state, err = a.client.Status()

		if err = a.LoadOppShots(state); err != nil {
			return "", err
		}

		printBoard(a.board, *a.desc)
		if err != nil {
			return "", err
		}

		if state.GameStatus == GameOver {
			break
		}

		if state.ShouldFire {
			a.Fire()
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

	fire, err := a.client.Fire(shot)
	if err != nil {
		return
	}

	switch fire {
	case "miss":
		a.board.Set(gui.Right, shot, gui.Miss)
	case "hit", "sunk":
		a.board.Set(gui.Right, shot, gui.Hit)
		if fire == "sunk" {
			a.board.CreateBorder(gui.Right, shot)
		}
	}
}

func (a *App) WaitForStart() (*Status, error) {
	for retryCount := 0; retryCount < maxRetries; retryCount++ {
		state, err := a.client.Status()
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
