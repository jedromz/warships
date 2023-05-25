package display

import (
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"time"
)

const (
	playerRows      = 1
	playerColumns   = 4
	opponentRows    = 50
	opponentColumns = 4
	defaultTimer    = 60
)

const (
	playerTurn   = "player_turn"
	opponentTurn = "opponent_turn"
	timer        = "timer"
)

const (
	ended            = "ended"
	game_in_progress = "game_in_progress"
)

type Game struct {
	gui           *gui.GUI
	playerBoard   *gui.Board
	opponentBoard *gui.Board
	s             *GameService
	dsc           Description
	gameChan      chan GameEvent
	resetTimer    chan bool
	timer         int
}

func NewApp() *Game {
	return &Game{
		gui:           gui.NewGUI(true),
		playerBoard:   gui.NewBoard(1, 4, nil),
		opponentBoard: gui.NewBoard(50, 4, nil),
		s:             NewGameService(),
		gameChan:      make(chan GameEvent),
		resetTimer:    make(chan bool),
		timer:         60,
	}
}

func (a *Game) StartBotGame() {
	for {
		fmt.Println("Would you like to place your ships manually?(y/n)")
		var answer string
		fmt.Scanln(&answer)

		var ships []string
		if answer == "y" {
			ships = a.placeShips()
		}
		fmt.Println(ships)
		a.gui.Log("StartAgain")
		err := a.s.StartBotGame(ships)
		if err != nil {
			return
		}

		_, err = a.s.LoadPlayerBoard()

		if err != nil {
			return
		}
		err = a.s.WaitForGame()
		if err != nil {
			return
		}

		a.dsc, err = a.s.GetDescription()
		if err != nil {
			return
		}
		go a.gameLoop()
		go a.startTimer()
		a.display()
		//abort?
		fmt.Println("Would you like to abort the ship?(y/n)")
		fmt.Scanln(&answer)
		if answer == "y" {
			a.s.AbortGame()
		}
		fmt.Println("Would you like to play again?(y/n)")
		_, err = fmt.Scanln(&answer)
		if err != nil {
			return
		}
		if answer == "n" {
			break
		}
	}
	return
}

func (a *Game) gameLoop() {
	var sts GameStatusResponse
	var err error
	for sts, err = a.s.UpdateGameStatus(); sts.GameStatus != ended && err == nil; sts, err = a.s.UpdateGameStatus() {
		a.gui.Draw(gui.NewText(1, 1, "Enemy turn!", nil))
		a.gameChan <- GameEvent{
			Type: "opponent_turn",
			Data: sts.OppShots,
		}
		if sts.ShouldFire {
			a.gui.Draw(gui.NewText(1, 1, "Your turn!", nil))
			a.gameChan <- GameEvent{
				Type: "player_turn",
				Data: "Your turn",
			}
			<-a.gameChan
		}
	}
	a.gui.Draw(gui.NewText(1, 1, "Game ended! "+sts.LastGameStatus, nil))
}

func (g *Game) placeShips() []string {
	g.gui.Draw(g.playerBoard)
	c := make(chan []string)
	var shipPlacement []string
	go func() {
		states := [10][10]gui.State{}

		ships := map[int]int{
			1: 4,
			2: 3,
			3: 2,
			4: 1,
		}
		for k, v := range ships {
			var coords []string
			for i := 0; i < k; i++ {
				for j := 0; j < v; j++ {
					coord := g.playerBoard.Listen(context.Background())
					coords = g.SetState(gui.Ship, coord, &states, coords)
				}
				err := g.s.PlaceShip(coords)
				if err != nil {
					return
				}
				shipPlacement = append(shipPlacement, coords...)
				coords = nil
				g.opponentBoard.SetStates(*g.s.GetOpponentFields())
			}

		}
		g.s.UpdatePlayerStates(states)
		c <- shipPlacement
	}()
	g.gui.Start(context.Background(), nil)
	return <-c
}

func (g *Game) SetState(state gui.State, coord string, states *[10][10]gui.State, coords []string) []string {
	x, y := mapToState(coord)
	states[x][y] = state
	coords = append(coords, coord)
	g.playerBoard.SetStates(*states)
	return coords
}

func (a *Game) display() {
	a.gui.Draw(a.playerBoard)
	a.gui.Draw(a.opponentBoard)
	a.gui.Draw(newFlag())
	a.drawDescription()

	go func() {
		for {
			event := <-a.gameChan
			switch event.Type {
			case playerTurn:
				coords := a.opponentBoard.Listen(context.TODO())
				err := a.s.Fire(coords)
				if err != nil {
					return
				}
				a.opponentBoard.SetStates(*a.s.GetOpponentFields())
				a.gameChan <- GameEvent{
					Type: opponentTurn,
				}

				if err != nil {
					return
				}
				a.drawAccuracy()
				a.resetTimer <- true
			case opponentTurn:
				shots := event.Data.([]string)
				a.s.UpdatePlayerBoard(shots)
				a.playerBoard.SetStates(*a.s.GetPlayerStates())
				a.resetTimer <- true
			case timer:
				a.timer = 60
			}
		}
	}()
	a.gui.Start(context.TODO(), nil)
}

func (a *Game) drawAccuracy() {
	accuracy := fmt.Sprintf("%.2f", a.s.GetPlayerAccuracy())
	text := "Your accuracy: " + accuracy + "%"
	a.gui.Draw(gui.NewText(1, 2, text, nil))

}

func (a *Game) drawDescription() {

	a.gui.Draw(gui.NewText(1, 26, a.dsc.Nick, nil))
	a.gui.Draw(gui.NewText(1, 27, a.dsc.Desc, nil))
	a.gui.Draw(gui.NewText(50, 26, a.dsc.Opponent, nil))
	a.gui.Draw(gui.NewText(50, 27, a.dsc.OppDesc, nil))
}
func (a *Game) startTimer() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			a.timer--

			// Stop the timer at zero
			if a.timer <= 0 {
				return
			}
			a.gui.Draw(gui.NewText(1, 3, fmt.Sprintf("Timer: %d seconds", a.timer), nil))

		case <-a.resetTimer:
			a.timer = 60
		}
	}
}
func (g *Game) listPlayers() {
	players, err := g.s.ListPlayers()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, player := range players {
		fmt.Println(player)
	}
}

func (a *Game) StartPvpGame() {
	var nick string
	var desc string
	var targetNick string
	fmt.Println("Enter your nick:")
	fmt.Scanln(&nick)
	fmt.Println("Enter your description:")
	fmt.Scanln(&desc)
	fmt.Println("Enter your target nick:")
	fmt.Scanln(&targetNick)

	for {
		a.s.StartPvpGame(nick, desc, targetNick)

		_, err := a.s.LoadPlayerBoard()
		if err != nil {
			break
		}
		err = a.s.WaitForGame()
		if err != nil {
			break
		}

		a.dsc, err = a.s.GetDescription()
		if err != nil {
			break
		}
		go a.gameLoop()
		go a.startTimer()
		a.display()
		fmt.Println("Would you like to play again?(y/n)")
		var answer string
		_, err = fmt.Scanln(&answer)
		if err != nil {
			break
		}
		if answer == "n" {
			break
		}
	}
}
