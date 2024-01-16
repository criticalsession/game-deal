package cmd

type command struct {
	name        string
	description string
	function    func(args ...string) error
}

func buildCommands() map[string]command {
	return map[string]command{
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
	}
}
