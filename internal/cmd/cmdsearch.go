package cmd

import (
	"fmt"
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/internal/utils"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
	"github.com/rodaine/table"
)

func cmdSearch(config *api.Config, args ...string) {
	t := strings.Join(args, " ")
	c := color.New(color.FgGreen)
	c.Printf("Searching for: \"%s\"\n", t)

	games, err := config.SearchGames(t)
	if err != nil {
		color.Red("%sAn error occured while searching games: %s", emoji.Sprintf(":red_exclamation_mark:"), err.Error())
		return
	}

	if len(games) == 0 {
		c := color.New(color.Reset)
		c.Printf("No games found :(\n")
		return
	}

	fmt.Println()

	headerFmt := color.New(color.FgGreen, color.Bold, color.Underline).SprintfFunc()
	idFmt := color.New(color.FgCyan).SprintfFunc()

	tbl := table.New("GameID", "| Title", "| Cheapest Deal")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(idFmt)

	for _, game := range games {
		sPrice, _ := utils.StringTo2fString(game.Cheapest)
		sPrice = "$" + sPrice

		tbl.AddRow(game.GameID, "| "+game.Title, "| "+sPrice)
	}

	tbl.Print()

	c = color.New(color.Reset)
	c.Println()
	c.Printf("Use \"")
	c = c.Add(color.Bold)
	c.Print("deals ")
	c = c.Add(color.FgHiCyan)
	c.Printf("[gameID]")
	c = color.New(color.Reset)
	c.Printf("\" command to see deals for a game\n")
}
