package cmd

import (
	"time"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/utils"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

func cmdOpenDeal(config *api.Config, args ...string) {
	if len(args) != 1 {
		utils.PrintError("\"deals\" command requires a gameID. Use \"search [keywords]\" command to find gameIDs.")
		return
	}

	index, err := utils.GetIndexFromInput(args[0])
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	deal, err := config.GetDealFromList(index)
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	color.Green("Opening deal in browser %s\n\n", emoji.Sprint(":rocket:"))

	dealUrl := config.GetDealUrl(deal.DealID)
	go sleepThenOpen(dealUrl)
}

func sleepThenOpen(url string) {
	time.Sleep(1 * time.Second)
	utils.OpenURL(url)
}
