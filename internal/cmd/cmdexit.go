package cmd

import "os"

func cmdExit(args ...string) error {
	os.Exit(0)
	return nil
}
