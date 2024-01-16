package api

import (
	"github.com/criticalsession/game-deal/internal/types/stores"
)

func (c *Config) GetStores() (stores.Resp, error) {
	url := baseUrl + "/stores"
	storeResp := stores.Resp{}

	err := getData(url, c.client, &storeResp)
	if err != nil {
		return stores.Resp{}, err
	}

	return stores.Resp(storeResp), nil
}
