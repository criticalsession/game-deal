package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/criticalsession/game-deal/utils"
	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

func CmdLoop(cfg *api.Config) {
	head()
	commands, _ := getCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := scanner.Text()

		args := cleanInput(input)
		if len(args) == 0 {
			printHelp(cfg, &commands)
			continue
		}

		cmd, ok := commands[args[0]]
		if !ok {
			printHelp(cfg, &commands)
			continue
		}

		cmd.function(cfg, args[1:]...)
	}
}

func head() {
	c := color.New(color.FgHiCyan).Add(color.Bold)
	c.Printf("\n%sGAME-DEAL %s", emoji.Sprintf(":video_game:"), emoji.Sprintf(":video_game:"))
	c = color.New(color.Reset).Add(color.Bold)
	c.Printf("\nFind and track game deals\n\n")

	c = color.New(color.Faint)
	c.Printf("This tool uses the CheapShark API to find and track game deals but is not affiliated with CheapShark in any way. ")
	c.Printf("Deal links use cheapshark's redirect url but should not affect your price.\n\n")

	utils.HelperLine("help", "", "get a list of available commands")
}

func cleanInput(s string) []string {
	s = strings.ToLower(s)
	return strings.Fields(s)
}

func printHelp(cfg *api.Config, commands *map[string]command) {
	(*commands)["help"].function(cfg)
}
