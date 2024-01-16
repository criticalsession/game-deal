package cmd

import (
	"github.com/criticalsession/game-deal/internal/api"
	"github.com/fatih/color"
)

func cmdDeals(config *api.Config, args ...string) {
	if len(args) != 1 {
		color.Red("\"deals\" command requires a gameID")
		return
	}
}
