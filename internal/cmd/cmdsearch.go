package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/utils"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
)

func cmdSearch(config *api.Config, args ...string) {
	titles := strings.Join(args, " ")
	c := color.New(color.FgGreen)
	c.Printf("Searching for: \"%s\"\n", titles)

	games, err := config.SearchGames(titles)
	if err != nil {
		utils.PrintError(fmt.Sprint("An error occured while searching games:", err.Error()))
		return
	}

	if len(games) == 0 {
		c := color.New(color.Reset)
		c.Printf("No games found :(\n")
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

	for _, game := range games {
		sPrice, _ := utils.StringTo2fString(game.Cheapest)
		sPrice = "$" + sPrice

		t.AppendRow([]interface{}{cyanFmt("[" + fmt.Sprint(gameIdCount+1) + "]"), game.Title,
			yellowFmt(sPrice)})

		gameIdCount++
	}

	t.SetStyle(table.StyleLight)
	t.Render()

	c = color.New(color.Reset)
	c.Println()
	c.Printf("Use \"")
	c = c.Add(color.Bold)
	c.Print("deals ")
	c = c.Add(color.FgHiCyan)
	c.Printf("[ID]")
	c = color.New(color.Reset)
	c.Printf("\" command to see deals for a game\n")

	config.SetGameList(games)
}
