package service

type GameStatusResponse struct {
	GameStatus     string   `json:"game_status"`
	LastGameStatus string   `json:"last_game_status"`
	Nick           string   `json:"nick"`
	OppShots       []string `json:"opp_shots"`
	Opponent       string   `json:"opponent"`
	ShouldFire     bool     `json:"should_fire"`
	Timer          int      `json:"timer"`
}
type GameBoard struct {
	Board []string `json:"board"`
}
type FireResponse struct {
	Coord string `json:"coord"`
}

type Description struct {
	Desc     string `json:"desc"`
	Nick     string `json:"nick"`
	OppDesc  string `json:"opp_desc"`
	Opponent string `json:"opponent"`
}

type LobbyEntry struct {
	GameStatus string `json:"game_status"`
	Nick       string `json:"nick"`
}

type GameEvent struct {
	Type string
	Data interface{}
}
type Player struct {
	GameStatus string `json:"game_status"`
	Nick       string `json:"nick"`
}
type GameStat struct {
	Games  int    `json:"games"`
	Nick   string `json:"nick"`
	Points int    `json:"points"`
	Rank   int    `json:"rank"`
	Wins   int    `json:"wins"`
}

type GameStats []GameStat

func (g GameStats) Len() int {
	return len(g)
}

func (g GameStats) Less(i, j int) bool {
	if g[i].Wins != g[j].Wins {
		return g[i].Wins > g[j].Wins
	}
	return g[i].Nick < g[j].Nick
}

func (g GameStats) Swap(i, j int) {
	g[i], g[j] = g[j], g[i]
}
