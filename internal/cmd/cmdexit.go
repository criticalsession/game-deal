package cmd

import (
	"os"

	"github.com/criticalsession/game-deal/internal/api"
)

func cmdExit(config *api.Config, args ...string) {
	os.Exit(0)
}
