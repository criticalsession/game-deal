package cmd

import "fmt"

func cmdPrintHelp(args ...string) error {
	fmt.Println("HELP!")
	return nil
}
