package cmd

import (
	"fmt"

	"github.com/criticalsession/game-deal/internal/api"
)

func cmdPrintHelp(config *api.Config, args ...string) {
	fmt.Println("Usage: > command [args]")
	fmt.Println()

	commands, order := getCommands()
	for _, cmd := range order {
		fmt.Printf("%s\n  - %s\n", commands[cmd].name, commands[cmd].description)
	}
}
