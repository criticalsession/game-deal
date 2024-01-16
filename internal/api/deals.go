package api

import (
	"strings"

	"github.com/criticalsession/game-deal/internal/types/gamedeals"
)

func (c *Config) GetGameDeals(ids ...string) (gamedeals.Resp, error) {
	url := baseUrl + "/games?ids=" + strings.Join(ids, ",")
	dealResp := gamedeals.Resp{}

	err := getData(url, c.client, &dealResp)
	if err != nil {
		return gamedeals.Resp{}, err
	}

	return gamedeals.Resp(dealResp), nil
}
