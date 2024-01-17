package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/db"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/kyokomi/emoji/v2"
)

func cmdFav(config *api.Config, args ...string) {
	if len(args) != 1 {
		color.Red("\"fav-add\" command requires a gameID. Use \"search [keywords]\" command to find gameIDs.")
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

	cheapsharkGame, err := config.GetGameFromGameList(id - 1)
	if err != nil {
		color.Red("%s%s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	err = db.AddFav(cheapsharkGame.GameID, cheapsharkGame.Title)
	if err != nil {
		color.Red("%s%s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	color.Green("\"%s\" successfully added to favorites!", cheapsharkGame.Title)
	fmt.Println()

	cmdListFav(config)
}

func cmdUnfav(config *api.Config, args ...string) {
	if len(args) != 1 {
		color.Red("\"fav-remove\" command requires a gameID. Use \"fav-list\" command to see favorited games.")
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

	game, err := db.GetFavByIndex(id - 1)
	if err != nil {
		color.Red("%s%s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	err = db.RemoveFav(game.Id)
	if err != nil {
		color.Red("%s%s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	color.Green("\"%s\" successfully removed from favorites", game.Title)

	fmt.Println()
	cmdListFav(config)
}

func cmdListFav(config *api.Config, args ...string) {
	res, err := db.GetFavs()
	if err != nil {
		color.Red("%s%s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
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
