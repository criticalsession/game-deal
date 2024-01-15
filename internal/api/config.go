package api

import (
	"net/http"
	"time"
)

type Config struct {
	DealUrl string
	BaseApi string
	Client  http.Client
}

type StoreCache struct {
	StoreID   int
	StoreName string
	Expires   time.Time
}

func NewConfig() *Config {
	return &Config{
		DealUrl: "https://www.cheapshark.com/redirect?dealID=",
		BaseApi: "https://www.cheapshark.com/api/1.0",
		Client: http.Client{
			Timeout: time.Minute,
		},
	}
}

func (c *Config) GetDealUrl(dealId string) string {
	return c.DealUrl + dealId
}
