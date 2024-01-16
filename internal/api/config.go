package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/criticalsession/game-deal/internal/types/gamedeals"
	"github.com/criticalsession/game-deal/internal/types/gamesearch"
)

const baseUrl = "https://www.cheapshark.com/api/1.0"
const dealUrl = "https://www.cheapshark.com/redirect?dealID="

type Config struct {
	client     http.Client
	storeCache *storeCache
	dealList   []gamedeals.Deal
	gameList   []gamesearch.GameInfo
}

func NewConfig() *Config {
	return &Config{
		client: http.Client{
			Timeout: time.Minute,
		},
		storeCache: NewStoreCache(),
		dealList:   []gamedeals.Deal{},
	}
}

func (c *Config) GetDealUrl(dealId string) string {
	return dealUrl + dealId
}

func (c *Config) SetDealsList(deals []gamedeals.Deal) {
	c.dealList = deals
}

func (c *Config) GetDealFromList(index int) (gamedeals.Deal, error) {
	if index >= len(c.dealList) {
		return gamedeals.Deal{}, errors.New("deal id " + fmt.Sprintf("[%d]", index+1) + " not found")
	}

	return c.dealList[index], nil
}

func (c *Config) SetGameList(games []gamesearch.GameInfo) {
	c.gameList = games
}

func (c *Config) GetGameIdFromGameList(index int) (string, error) {
	if index >= len(c.gameList) {
		return "", errors.New("game id " + fmt.Sprintf("[%d]", index+1) + " not found")
	}

	return c.gameList[index].GameID, nil
}
