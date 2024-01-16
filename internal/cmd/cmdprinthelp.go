package cmd

import (
	"fmt"

	"github.com/criticalsession/game-deal/internal/api"
)

func cmdPrintHelp(config *api.Config, args ...string) {
	fmt.Println("Usage: > command [args]")
	fmt.Println()

	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("%s\n  - %s\n", cmd.name, cmd.description)
	}
}
