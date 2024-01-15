package api

import (
	"github.com/criticalsession/game-deal/internal/types/stores"
)

func (c *Config) GetStores() (stores.Resp, error) {
	url := c.BaseApi + "/stores"
	storeResp := stores.Resp{}

	err := getData(url, c.Client, &storeResp)
	if err != nil {
		return stores.Resp{}, err
	}

	return stores.Resp(storeResp), nil
}
