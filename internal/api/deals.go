package api

import (
	"github.com/criticalsession/game-deal/internal/types/gamedeals"
)

func (c *Config) GetGameDeals(id string) (gamedeals.Resp, error) {
	url := baseUrl + "/games?ids=" + id
	dealResp := gamedeals.Resp{}

	err := getData(url, c.client, &dealResp)
	if err != nil {
		return gamedeals.Resp{}, err
	}

	return gamedeals.Resp(dealResp), nil
}
