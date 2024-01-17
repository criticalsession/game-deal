package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/types/gamedeals"
	"github.com/criticalsession/game-deal/utils"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func cmdDeals(config *api.Config, args ...string) {
	if len(args) != 1 {
		utils.PrintError("\"deals\" command requires a gameID. Use \"search [keywords]\" command to find gameIDs.")
		return
	}

	index, err := utils.GetIndexFromInput(args[0])
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	cheapsharkGame, err := config.GetGameFromGameList(index)
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	result, err := config.GetGameDeals(cheapsharkGame.GameID)
	if err != nil {
		utils.PrintError(fmt.Sprint("An error occured while searching deals:", err.Error()))
		return
	}

	stores, err := config.GetStores()
	if err != nil {
		utils.PrintError(fmt.Sprint("An error occured while getting stores:", err.Error()))
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
