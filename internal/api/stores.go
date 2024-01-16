package api

import (
	"strconv"

	"github.com/criticalsession/game-deal/internal/types/stores"
)

func (c *Config) GetStores() (map[int]stores.Store, error) {
	url := baseUrl + "/stores"
	storeResp := stores.Resp{}

	err := getData(url, c.client, &storeResp)
	if err != nil {
		return map[int]stores.Store{}, err
	}

	result := make(map[int]stores.Store)
	for _, store := range storeResp {
		id, err := strconv.Atoi(store.StoreID)
		if err != nil {
			return map[int]stores.Store{}, err
		}

		result[id] = store
	}

	return result, nil
}
