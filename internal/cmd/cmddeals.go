package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/types/gamedeals"
	"github.com/criticalsession/game-deal/utils"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kyokomi/emoji/v2"
)

func cmdDeals(config *api.Config, args ...string) {
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

	cheapsharkGame, err := config.GetGameFromGameList(id)
	if err != nil {
		color.Red("%s%s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	result, err := config.GetGameDeals(cheapsharkGame.GameID)
	if err != nil {
		color.Red("%sAn error occured while searching deals: %s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	stores, err := config.GetStores()
	if err != nil {
		color.Red("%sAn error occured while getting stores: %s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	dealList := []gamedeals.Deal{}
	dealListCount := 0

	headerFmt := color.New(color.FgGreen, color.Bold).SprintFunc()
	cyanFmt := color.New(color.FgCyan).SprintFunc()
	greenFmt := color.New(color.FgGreen).SprintFunc()
	yellowFmt := color.New(color.FgYellow).SprintFunc()

	for _, res := range result {
		c := color.New(color.FgGreen)
		c.Printf("Deals for: \"%s\"\n", res.Info.Title)

		c = color.New(color.FgRed)
		sCheapest, _ := utils.StringTo2fString(res.CheapestEver.Price)
		c.Printf("Historically cheapest price: $%s (%s)\n", sCheapest, utils.UnixToDateString(int64(res.CheapestEver.Date)))

		fmt.Println()

		if len(res.Deals) == 0 {
			c := color.New(color.Reset)
			c.Printf("No active deals found :(\n")
			continue
		}

		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{headerFmt("ID"), headerFmt("Store"), headerFmt("Original Price"), headerFmt("Discounted Price"), headerFmt("Savings")})

		for _, deal := range res.Deals {
			sPrice, _ := utils.StringTo2fString(deal.Price)
			sRetailPrice, _ := utils.StringTo2fString(deal.RetailPrice)
			sSavings, _ := utils.StringTo2fString(deal.Savings)
			iStoreID, _ := strconv.Atoi(deal.StoreID)

			s, _ := strconv.ParseFloat(sSavings, 64)
			if s > 0 {
				sSavings = greenFmt(sSavings + "%")
			} else {
				sSavings += "%"
			}

			t.AppendRow([]interface{}{cyanFmt("[" + fmt.Sprint(dealListCount+1) + "]"), stores[iStoreID].StoreName,
				yellowFmt("$" + sRetailPrice), yellowFmt("$" + sPrice), sSavings})

			dealList = append(dealList, deal)
			dealListCount++
		}

		t.SetStyle(table.StyleLight)
		t.Render()

		fmt.Println("")
	}

	if len(dealList) > 0 {
		c := color.New(color.Reset)
		c.Printf("Use \"")
		c = c.Add(color.Bold)
		c.Print("open ")
		c = c.Add(color.FgHiCyan)
		c.Printf("[ID]")
		c = color.New(color.Reset)
		c.Printf("\" command to open deal in browser\n")
	}

	config.SetDealsList(dealList)
}
