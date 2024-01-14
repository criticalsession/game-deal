package gamedeals

type Resp map[string]SingleGameResp

type SingleGameResp struct {
	Info         GameInfo      `json:"info"`
	CheapestEver CheapestPrice `json:"cheapestPriceEver"`
	Deals        []Deal        `json:"deals"`
}

type GameInfo struct {
	Title      string `json:"title"`
	SteamAppID any    `json:"steamAppID"`
}

type CheapestPrice struct {
	Price string `json:"price"`
	Date  int    `json:"date"`
}

type Deal struct {
	StoreID     string `json:"storeID"`
	DealID      string `json:"dealID"`
	Price       string `json:"price"`
	RetailPrice string `json:"retailPrice"`
	Savings     string `json:"savings"`
}
