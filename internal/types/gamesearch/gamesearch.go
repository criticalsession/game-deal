package gamesearch

type Resp []GameInfo

type GameInfo struct {
	GameID         string `json:"gameID"`
	SteamAppID     string `json:"steamAppID"`
	Cheapest       string `json:"cheapest"`
	CheapestDealID string `json:"cheapestDealID"`
	Title          string `json:"external"`
	Thumb          string `json:"thumb"`
}
