package cmd

import (
	"github.com/criticalsession/game-deal/internal/api"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

func cmdStores(config *api.Config, args ...string) {
	stores, err := config.GetStores()
	if err != nil {
		color.Red("%sAn error occured while getting stores: %s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	for _, s := range stores {
		c := color.New(color.FgHiCyan)
		c.Printf("%-6s", "["+s.StoreID+"]")
		c = color.New(color.Reset).Add(color.Bold)
		c.Printf("%s\n", s.StoreName)

	}
}
