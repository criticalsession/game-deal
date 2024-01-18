package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/db"
	"github.com/criticalsession/game-deal/utils"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func cmdFav(config *api.Config, args ...string) {
	if len(args) != 1 {
		utils.PrintError("\"fav-add\" command requires a gameID. Use \"search [keywords]\" command to find gameIDs.")
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

	err = db.AddFav(cheapsharkGame.GameID, cheapsharkGame.Title)
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	color.Green("\"%s\" successfully added to favorites!", cheapsharkGame.Title)
	fmt.Println()

	cmdListFav(config)
}

func cmdUnfav(config *api.Config, args ...string) {
	if len(args) != 1 {
		utils.PrintError("\"fav-remove\" command requires a gameID. Use \"fav-list\" command to see favorited games.")
		return
	}

	index, err := utils.GetIndexFromInput(args[0])
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	game, err := db.GetFavByIndex(index)
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	err = db.RemoveFav(game.Id)
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	color.Green("\"%s\" successfully removed from favorites", game.Title)

	fmt.Println()
	cmdListFav(config)
}

func cmdListFav(config *api.Config, args ...string) {
	res, err := db.GetFavs()
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	if len(res) == 0 {
		color.Yellow("No favorited games :(\n\n")
		return
	}

	headerFmt := color.New(color.FgGreen, color.Bold).SprintFunc()
	cyanFmt := color.New(color.FgCyan).SprintfFunc()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{headerFmt("ID"), headerFmt("Title")})

	for i, r := range res {
		t.AppendRow([]interface{}{cyanFmt("[" + fmt.Sprint(i+1) + "]"), r.Title})
	}

	t.SetStyle(table.StyleLight)
	t.Render()

	fmt.Println()
}

func cmdFavDeals(config *api.Config, args ...string) {
	res, err := db.GetFavs()
	if err != nil {
		utils.PrintError(err.Error())
		return
	}

	if len(res) == 0 {
		color.Yellow("No favorited games :(\n\n")
		return
	}

	ids := []string{}

	for _, r := range res {
		ids = append(ids, r.GameId)
	}

	getGameDeals(config, strings.Join(ids, ","))
}
