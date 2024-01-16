package cmd

import (
	"strconv"
	"strings"
	"time"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/utils"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

func cmdOpenDeal(config *api.Config, args ...string) {
	if len(args) != 1 {
		color.Red("\"deals\" command requires a gameID. Use \"search [keywords]\" command to find gameIDs.")
		return
	}

	sid := args[0]

	sid = strings.ReplaceAll(sid, "[", "")
	sid = strings.ReplaceAll(sid, "]", "")

	id, err := strconv.Atoi(sid)
	if err != nil {
		color.Red("%sInvalid id: %s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}
	id -= 1

	deal, err := config.GetDealFromList(id)
	if err != nil {
		color.Red("%s%s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
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
