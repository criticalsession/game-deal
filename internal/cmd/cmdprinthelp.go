package cmd

import (
	"fmt"
	"strings"

	"github.com/criticalsession/game-deal/internal/api"
	"github.com/fatih/color"
)

func cmdPrintHelp(config *api.Config, args ...string) {
	c := color.New(color.Reset).Add(color.Bold)
	c.Println("Usage: > command [args] <optional>")
	fmt.Println()
	c.Println("Commands:")

	boldFmt := color.New(color.Reset).Add(color.Bold).SprintFunc()

	commands, order := getCommands()
	for _, cmd := range order {
		fmt.Printf("  %s\n    - %s\n", boldFmt(commands[cmd].name), strings.ReplaceAll(commands[cmd].description, "\n", "\n    - "))
	}
}
