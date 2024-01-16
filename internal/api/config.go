package api

import (
	"net/http"
	"time"
)

const baseUrl = "https://www.cheapshark.com/api/1.0"
const dealUrl = "https://www.cheapshark.com/redirect?dealID="

type Config struct {
	client     http.Client
	storeCache *storeCache
}

func NewConfig() *Config {
	return &Config{
		client: http.Client{
			Timeout: time.Minute,
		},
		storeCache: NewStoreCache(),
	}
}

func (c *Config) GetDealUrl(dealId string) string {
	return dealUrl + dealId
}
