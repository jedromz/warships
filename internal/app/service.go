package app

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/fatih/color"
	gui "github.com/grupawp/warships-lightgui/v2"
	"log"
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

func (a *App) Play() error {
	for {
		state, err := a.client.Status()
		fmt.Println(state)
		//load initial opponent shots
		if err = a.LoadOppShots(state); err != nil {
			return err
		}
		printBoard(a.board, state)

		if err != nil {
			return err
		}

		if state.GameStatus == GameOver {
			break
		}

		if state.ShouldFire {
			a.Fire()
		}
		printBoard(a.board, state)
	}

	return nil
}

func (a *App) LoadOppShots(state *Status) error {
	a.oppShots = state.OppShots[len(a.oppShots):]
	for _, shot := range a.oppShots {
		mark, err := a.board.HitOrMiss(gui.Left, shot)
		if err != nil {
			return err
		}

		a.board.Set(gui.Left, shot, mark)
	}

	return nil
}

func (a *App) Fire() {
	shot := enterCords()
	if shot == "stop" {
		os.Exit(0)
	}

	fire, err := a.client.Fire(shot)
	if err != nil {
		fmt.Println(err)
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

func printBoard(board *gui.Board, status *Status) {
	opponent := color.New(color.FgHiWhite, color.BgRed).SprintFunc()
	player := color.New(color.FgHiWhite, color.BgGreen).SprintFunc()

	board.Display()

	playerName := getOrDefault(status.Desc, "Player")
	opponentName := getOrDefault(status.OppDesc, "Flying WP Bot")

	fmt.Printf("%-30s%s%30s\n", player(status.Nick), "", opponent(status.Opponent))
	fmt.Printf("%-37s%s%30s\n", player(playerName), "", opponent(opponentName))
	fmt.Println()

}

func enterCords() string {
	fmt.Print("Enter shot: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}
	return strings.TrimSpace(text)
}

func getOrDefault(s, d string) string {
	if s == "" {
		return d
	}
	return s
}
