package main

import (
	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/cmd"
	"github.com/criticalsession/game-deal/internal/db"
)

func main() {
	cfg := api.NewConfig()
	_, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	cmd.CmdLoop(cfg)
}
