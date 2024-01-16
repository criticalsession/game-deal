package cmd

import (
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
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

	for _, game := range games {
		c := color.New(color.FgCyan)
		c.Printf("%-11s", "["+game.GameID+"]")
		c = color.New(color.Reset).Add(color.Bold)
		c.Printf("%s\n", game.Title)
	}

	c = color.New(color.Reset)
	c.Println()
	c.Printf("Use \"")
	c = c.Add(color.Bold)
	c.Print("deals ")
	c = c.Add(color.FgCyan)
	c.Printf("[gameID]")
	c = color.New(color.Reset)
	c.Printf("\" command to see deals for a game\n")
}
