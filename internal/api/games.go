package api

import (
	"net/url"

	"github.com/criticalsession/game-deal/internal/types/gamesearch"
)

func (c *Config) SearchGames(title string) (gamesearch.Resp, error) {
	url := baseUrl + "/games?title=" + url.QueryEscape(title)
	gameResp := gamesearch.Resp{}

	err := getData(url, c.client, &gameResp)
	if err != nil {
		return gamesearch.Resp{}, err
	}

	return gamesearch.Resp(gameResp), nil
}
