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
