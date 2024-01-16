package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/criticalsession/game-deal/internal/api"
)

func main() {
	cfg := api.NewConfig()
	res, err := cfg.GetGameDeals("128")
	if err != nil {
		log.Fatal(err)
	}

	stores, err := cfg.GetStores()
	if err != nil {
		log.Fatal(err)
	}

	for _, x := range res {
		fmt.Printf("Cheapest Historical Price: $%s\n", x.CheapestEver.Price)
		i := 0
		for _, deal := range x.Deals {
			storeId, _ := strconv.Atoi(deal.StoreID)
			savings, _ := strconv.ParseFloat(deal.Savings, 32)
			i++
			if savings > 0 {
				fmt.Printf("[%d] @%s: $%s > $%s (%.1f%% savings)\n",
					i, stores[storeId].StoreName, deal.RetailPrice, deal.Price, savings)
			}
		}
	}

	openURL(cfg.GetDealUrl(res["128"].Deals[0].DealID))
}

// https://apidocs.cheapshark.com/
