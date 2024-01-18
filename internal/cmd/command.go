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
				name:        "deals [gameID] <max=?>",
				description: "find deals for given game\nmax is optional and limits maximum price (e.g. max=30)",
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
			"fav-list": {
				name:        "fav-list",
				description: "list all favorited games",
				function:    cmdListFav,
			},
			"fav-add": {
				name:        "fav-add [gameID]",
				description: "add game to favorites",
				function:    cmdFav,
			},
			"fav-remove": {
				name:        "fav-remove [gameID]",
				description: "remove game from favorites",
				function:    cmdUnfav,
			},
			"fav-deals": {
				name:        "fav-deals <max=?>",
				description: "get deals for all your favorited games\nmax is optional and limits maximum price (e.g. max=30)",
				function:    cmdFavDeals,
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
		}, []string{"search", "deals", "open", "stores", "fav-list",
			"fav-add", "fav-remove", "fav-deals", "help", "exit"} // correct order for help
}
