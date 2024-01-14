package config

type Config struct {
	DealUrl string
}

func New() *Config {
	return &Config{
		DealUrl: "https://www.cheapshark.com/redirect?dealID=",
	}
}

func (c *Config) GetDealUrl(dealId string) string {
	return c.DealUrl + dealId
}
