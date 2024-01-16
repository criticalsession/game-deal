package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/types/gamedeals"
	"github.com/criticalsession/game-deal/internal/utils"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
	"github.com/rodaine/table"
)

func cmdDeals(config *api.Config, args ...string) {
	if len(args) != 1 {
		color.Red("\"deals\" command requires a gameID. Use \"search [keywords]\" command to find gameIDs.")
		return
	}

	id := args[0]

	id = strings.ReplaceAll(id, "[", "")
	id = strings.ReplaceAll(id, "]", "")

	result, err := config.GetGameDeals(id)
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

	headerFmt := color.New(color.FgGreen, color.Bold, color.Underline).SprintfFunc()
	idFmt := color.New(color.FgCyan).SprintfFunc()

	for _, res := range result {
		c := color.New(color.FgGreen)
		c.Printf("Deals for: \"%s\"\n", res.Info.Title)

		c = color.New(color.FgRed)
		sCheapest, _ := utils.StringTo2fString(res.CheapestEver.Price)
		c.Printf("Historically cheapest price: $%s (%s)\n", sCheapest, utils.UnixToDateString(int64(res.CheapestEver.Date)))

		fmt.Println()

		savingDeals := []gamedeals.Deal{}
		for _, deal := range res.Deals {
			savings, err := strconv.ParseFloat(deal.Savings, 64)
			if err != nil {
				savings = 0
			}

			if savings > 0 {
				savingDeals = append(savingDeals, deal)
			}
		}

		if len(savingDeals) == 0 {
			c := color.New(color.Reset)
			c.Printf("No active deals found :(\n")
			continue
		}

		tbl := table.New("DealID", "| Store", "| Original Price", "| Discounted Price", "| Savings")
		tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(idFmt)

		for _, deal := range savingDeals {
			sPrice, _ := utils.StringTo2fString(deal.Price)
			sRetailPrice, _ := utils.StringTo2fString(deal.RetailPrice)
			sSavings, _ := utils.StringTo2fString(deal.Savings)
			iStoreID, _ := strconv.Atoi(deal.StoreID)

			tbl.AddRow(dealListCount+1, "| "+stores[iStoreID].StoreName, "| $"+sRetailPrice, "| $"+sPrice, "| "+sSavings+"%")

			dealList = append(dealList, deal)
			dealListCount++
		}

		tbl.Print()

		fmt.Println("")
	}

	if len(dealList) > 0 {
		c := color.New(color.Reset)
		c.Printf("Use \"")
		c = c.Add(color.Bold)
		c.Print("open ")
		c = c.Add(color.FgHiCyan)
		c.Printf("[dealID]")
		c = color.New(color.Reset)
		c.Printf("\" command to open deal in browser\n")
	}

	config.SetDealsList(dealList)
}
