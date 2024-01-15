package main

import (
	"log"

	"github.com/criticalsession/game-deal/internal/api"
)

func main() {
	cfg := api.NewConfig()
	res, err := cfg.GetGameDeals("144222")
	if err != nil {
		log.Fatal(err)
	}

	for _, x := range res {
		log.Println(x)
	}
}

// https://apidocs.cheapshark.com/
