package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
)

func CmdLoop() {
	head()

	for {
		fmt.Print("> ")
		var input string
		fmt.Scanln(&input)

		args := cleanInput(input)
		if len(args) == 0 {
			printHelp()
			continue
		}

		if args[0] == "exit" {
			os.Exit(0)
		}
	}
}

func head() {
	c := color.New(color.FgHiCyan).Add(color.Bold)
	c.Printf("\n%s GAME-DEAL %s", emoji.Sprintf(":video_game:"), emoji.Sprintf(":video_game:"))
	c = color.New(color.Reset).Add(color.Bold)
	c.Printf("\nFind and track game deals\n\n")

	c = color.New(color.Faint)
	c.Printf("This tool uses the cheapshark api to find and track game deals but is not affiliate with cheapshark in any way. ")
	c.Printf("Deal links use cheapshark's redirect url but does not affect your price.\n\n")

	c = color.New(color.Reset)
	c.Printf("Try 'help' for list of available commands.\n")
}

func cleanInput(s string) []string {
	s = strings.ToLower(s)
	return strings.Fields(s)
}

func printHelp() {
	fmt.Println("HELP!")
}
