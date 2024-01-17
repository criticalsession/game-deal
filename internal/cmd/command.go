package cmd

import "github.com/criticalsession/game-deal/internal/api"

type command struct {
	name        string
	description string
	function    func(config *api.Config, args ...string)
}

func getCommands() (map[string]command, []string) {
	return map[string]command{
		"search": {
			name:        "search [keywords]",
			description: "find games matching keywords",
			function:    cmdSearch,
		},
		"deals": {
			name:        "deals [gameID]",
			description: "find deals for given game",
			function:    cmdDeals,
		},
		"open": {
			name:        "open [dealID]",
			description: "open deal URL in browser",
			function:    cmdOpenDeal,
		},
		"stores": {
			name:        "stores",
			description: "list all stores available for deal search",
			function:    cmdStores,
		},
		"help": {
			name:        "help",
			description: "print this help message",
			function:    cmdPrintHelp,
		},
		"exit": {
			name:        "exit",
			description: "exit game-deal",
			function:    cmdExit,
		},
	}, []string{"search", "deals", "open", "stores", "help", "exit"} // correct order for help
}
