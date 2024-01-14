package stores

type Resp []Store

type Store struct {
	StoreID   string `json:"storeID"`
	StoreName string `json:"storeName"`
	IsActive  int    `json:"isActive"`
	Images    Images `json:"images"`
}

type Images struct {
	Banner string `json:"banner"`
	Logo   string `json:"logo"`
	Icon   string `json:"icon"`
}
