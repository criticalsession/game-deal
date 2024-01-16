package main

import (
	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/cmd"
)

func main() {
	cfg := api.NewConfig()
	cmd.CmdLoop(cfg)
}

// https://apidocs.cheapshark.com/
