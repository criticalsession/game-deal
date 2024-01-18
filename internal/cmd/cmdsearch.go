package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/types/gamesearch"
	"github.com/criticalsession/game-deal/utils"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func cmdSearch(config *api.Config, args ...string) {
	var keywords string
	var maxPrice float64

	if len(args) < 1 {
		utils.PrintError("\"search\" command requires at least one keyword.")
		return
	}

	for _, s := range args {
		if strings.Contains(s, "max=") {
			var err error
			maxPrice, err = getMaxPrice([]string{s}...)

			if err != nil {
				utils.PrintError(err.Error())
				return
			}

			continue
		}

		keywords += s + " "
	}

	c := color.New(color.FgGreen)
	c.Printf("Searching for: \"%s\"\n", keywords)

	games, err := config.SearchGames(keywords)
	if err != nil {
		utils.PrintError(fmt.Sprint("An error occured while searching games:", err.Error()))
		return
	}

	if len(games) == 0 {
		utils.PrintInfo("No games found")
		return
	}

	fmt.Println()

	headerFmt := color.New(color.FgGreen, color.Bold).SprintFunc()
	cyanFmt := color.New(color.FgCyan).SprintfFunc()
	yellowFmt := color.New(color.FgYellow).SprintfFunc()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{headerFmt("ID"), headerFmt("Title"), headerFmt("Cheapest Deal")})

	gameIdCount := 0

	filteredGames := gamesearch.Resp{}

	for _, game := range games {
		if maxPrice > 0 {
			fPriceFloat, _ := strconv.ParseFloat(game.Cheapest, 64)

			if fPriceFloat > maxPrice {
				continue
			}
		}

		sPrice, _ := utils.StringTo2fString(game.Cheapest)
		sPrice = "$" + sPrice

		t.AppendRow([]interface{}{cyanFmt("[" + fmt.Sprint(gameIdCount+1) + "]"), game.Title,
			yellowFmt(sPrice)})

		gameIdCount++

		filteredGames = append(filteredGames, game)
	}

	if gameIdCount == 0 && maxPrice > 0 {
		utils.PrintInfo(fmt.Sprintf("No games found with max price $%.2f", maxPrice))
		return
	}

	t.SetStyle(table.StyleLight)
	t.Render()

	fmt.Println()
	utils.HelperLine("deals", "[ID] <max=?>", "see deals for a game")

	config.SetGameList(filteredGames)
}
