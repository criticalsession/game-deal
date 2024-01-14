package main

import (
	"fmt"

	"github.com/criticalsession/game-deal/config"
)

func main() {
	cfg := config.New()
	fmt.Println(cfg.GetDealUrl("123"))
}

// https://apidocs.cheapshark.com/
