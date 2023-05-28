package game

import (
	"battleships/config"
	globals "battleships/globals"
	"battleships/service"
	"context"
	"fmt"
	gui "github.com/grupawp/warships-gui/v2"
	"sort"
	"time"
)

var (
	enemyTurnText  = gui.NewText(1, 1, "Enemy turn!", nil)
	playerTurnText = gui.NewText(1, 1, "Your turn!", nil)
)

const (
	playAgain  = "Would you like to play again?(y/n)"
	placeShips = "Would you like to place your ships manually?(y/n)"
	abort      = "Would you like to abort the game?(y/n)"
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
	s             *service.GameService
	dsc           globals.Description
	gameChan      chan globals.GameEvent
	resetTimer    chan bool
	timer         int
	gameActive    chan bool
}

func NewApp() *Game {
	return &Game{
		gui:           gui.NewGUI(true),
		playerBoard:   gui.NewBoard(config.PlayerBoardX, config.PlayerBoardY, nil),
		opponentBoard: gui.NewBoard(config.OpponentBoardX, config.OpponentBoardY, nil),
		s:             service.NewGameService(),
		gameChan:      make(chan globals.GameEvent),
		resetTimer:    make(chan bool),
		timer:         config.Timer,
		gameActive:    make(chan bool),
	}
}

func (g *Game) StartBotGame() {
	for {
		answer := askShipPlacement()

		var ships []string
		if answer == "y" {
			ships = g.placeShips()
		}

		err := g.s.StartBotGame(ships)
		if err != nil {
			g.gui.Log(err.Error())
		}

		if answer == "n" {
			_, err = g.s.LoadPlayerBoard()
		}
		if err != nil {
			return
		}

		go showSpinner(g.gameActive)
		err = g.s.WaitForGame()
		if err != nil {
			return

		}
		g.gameActive <- true
		g.dsc, err = g.s.GetDescription()
		if err != nil {
			return
		}

		go g.gameLoop()
		go g.startTimer()
		g.display()

		fmt.Println(abort)

		fmt.Scanln(&answer)
		if answer == "y" {
			g.s.AbortGame()
		}
		fmt.Println(playAgain)
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

func (g *Game) gameLoop() {
	var sts globals.GameStatusResponse
	var err error
	for sts, err = g.s.UpdateGameStatus(); sts.GameStatus != ended && err == nil; sts, err = g.s.UpdateGameStatus() {
		g.gui.Draw(enemyTurnText)
		g.gameChan <- globals.GameEvent{
			Type: opponentTurn,
			Data: sts.OppShots,
		}
		if sts.ShouldFire {
			g.gui.Draw(playerTurnText)
			g.gameChan <- globals.GameEvent{
				Type: playerTurn,
			}
			select {
			case <-g.gameChan:
				// Continue the game loop.
			case <-time.After(time.Duration(g.timer) * time.Second):
				g.gui.Draw(gui.NewText(25, 10, "Time's up!... Press ctrl c quit the game", nil))
				return
			}
		}
	}
	g.drawGameResult(sts)
}

var Ships = map[int]int{
	4: 1,
	3: 2,
	2: 3,
	1: 4,
}

func (g *Game) placeShips() []string {
	g.gui.Draw(g.playerBoard)
	c := make(chan []string)
	var shipPlacement []string

	go func() {
		states := [10][10]gui.State{}

		for shipLength, shipCount := range Ships {
			for i := 0; i < shipCount; i++ {
				var coords []string
				wellPlaced := false

				for !wellPlaced {
					coords = g.getValidShipPlacement(shipLength, &states)
					err := g.s.PlaceShip(coords)

					if err != nil {
						g.undoShipPlacement(coords, &states)
						coords = []string{}
					} else {
						shipPlacement = append(shipPlacement, coords...)
						g.opponentBoard.SetStates(*g.s.GetOpponentFields())
						g.s.UpdatePlayerStates(states)
						wellPlaced = true
					}
				}
			}
		}

		g.s.UpdatePlayerStates(states)
		c <- shipPlacement
	}()

	g.gui.Start(context.Background(), nil)
	return <-c
}

func (g *Game) getValidShipPlacement(shipLength int, states *[10][10]gui.State) []string {
	var coords []string

	for i := 0; i < shipLength; i++ {
		for {
			coord := g.playerBoard.Listen(context.Background())
			if !containsCoordinate(coord, coords) {
				coords = g.SetState(gui.Ship, coord, states, coords)
				break
			}
		}
	}

	return coords
}

func (g *Game) undoShipPlacement(coords []string, states *[10][10]gui.State) {
	for _, crd := range coords {
		coords = g.SetState(gui.Empty, crd, states, coords)
	}
}

func (g *Game) SetState(state gui.State, coord string, states *[10][10]gui.State, coords []string) []string {
	x, y := mapToState(coord)
	states[x][y] = state
	coords = append(coords, coord)
	g.playerBoard.SetStates(*states)
	return coords
}

func (g *Game) display() {
	g.gui.Draw(g.playerBoard)
	g.gui.Draw(g.opponentBoard)
	g.drawDescription()

	go func() {
		for {
			event := <-g.gameChan
			switch event.Type {
			case playerTurn:
				g.fire()
				g.gameChan <- globals.GameEvent{
					Type: opponentTurn,
				}
				g.drawAccuracy()
				g.resetClock()
			case opponentTurn:
				g.drawOpponentShots(event)
				g.resetClock()
			case timer:
				g.timer = config.Timer
			}
		}
	}()
	g.gui.Start(context.TODO(), nil)
}

func (g *Game) drawOpponentShots(event globals.GameEvent) {
	shots := get(event)
	g.s.UpdatePlayerBoard(shots)
	g.playerBoard.SetStates(*g.s.GetPlayerStates())
}

func (g *Game) resetClock() {
	g.resetTimer <- true
}

func get(event globals.GameEvent) []string {
	return event.Data.([]string)
}

func (g *Game) fire() {
	for {
		coords := g.opponentBoard.Listen(context.TODO())
		err := g.s.Fire(coords)
		if err != nil {
			g.gui.Draw(gui.NewText(1, 5, "You already fired there", nil))
		} else {
			break
		}
	}
	g.opponentBoard.SetStates(*g.s.GetOpponentFields())
}

func (g *Game) drawAccuracy() {
	accuracy := fmt.Sprintf("%.2f", g.s.GetPlayerAccuracy())
	text := "Your accuracy: " + accuracy + "%"
	g.gui.Draw(gui.NewText(1, 2, text, nil))

}

func (g *Game) drawDescription() {

	g.gui.Draw(gui.NewText(1, 26, g.dsc.Nick, nil))
	g.gui.Draw(gui.NewText(1, 27, g.dsc.Desc, nil))
	g.gui.Draw(gui.NewText(50, 26, g.dsc.Opponent, nil))
	g.gui.Draw(gui.NewText(50, 27, g.dsc.OppDesc, nil))
}
func (g *Game) startTimer() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			g.timer--
			if g.timer <= 0 {
				return
			}
			g.gui.Draw(gui.NewText(1, 3, fmt.Sprintf("Timer: %d seconds", g.timer), nil))

		case <-g.resetTimer:
			g.timer = config.Timer
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
		fmt.Printf("Game Status: %s\n", player.GameStatus)
		fmt.Printf("Nickname: %s\n", player.Nick)
		fmt.Println()
	}
}

func (g *Game) StartPvpGame() {
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
		answer := askShipPlacement()

		var ships []string
		if answer == "y" {
			ships = g.placeShips()
		}

		g.s.StartPvpGame(nick, desc, targetNick, ships)

		_, err := g.s.LoadPlayerBoard()
		if err != nil {
			break
		}

		go showSpinner(g.gameActive)
		err = g.s.WaitForGame()
		g.gameActive <- true

		g.dsc, err = g.s.GetDescription()
		if err != nil {
			break
		}

		go g.gameLoop()
		go g.startTimer()
		g.display()

		fmt.Println("Would you like to play again? (y/n)")

		_, err = fmt.Scanln(&answer)
		if err != nil {
			break
		}
		if answer == "n" {
			break
		}
	}
}

func (g *Game) ranking() {
	stats, err := g.s.GetStats()
	if err != nil {
		fmt.Println(err)
		return
	}
	sort.Sort(stats)

	for _, stat := range stats {
		fmt.Printf("Nickname: %s\n", stat.Nick)
		fmt.Printf("Games: %d\n", stat.Games)
		fmt.Printf("Points: %d\n", stat.Points)
		fmt.Printf("Rank: %d\n", stat.Rank)
		fmt.Printf("Wins: %d\n", stat.Wins)
		fmt.Println("---------------------------")
	}
}

func askShipPlacement() string {
	fmt.Println(placeShips)
	var answer string
	fmt.Scanln(&answer)
	return answer
}
func (g *Game) drawGameResult(sts globals.GameStatusResponse) {
	g.gui.Draw(gui.NewText(1, 1, "Game ended! "+sts.LastGameStatus, nil))
}
func containsCoordinate(coord string, coords []string) bool {
	for _, c := range coords {
		if c == coord {
			return true
		}
	}
	return false
}
