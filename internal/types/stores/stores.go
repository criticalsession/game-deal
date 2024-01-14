package stores

type Resp []Store

type Store struct {
	StoreID   string `json:"storeID"`
	StoreName string `json:"storeName"`
	IsActive  int    `json:"isActive"`
}
