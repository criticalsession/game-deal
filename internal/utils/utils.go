package utils

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
)

func StringTo2fString(price string) (string, error) {
	f, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f", f), nil
}

func OpenURL(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to open URL: %w", err)
	}

	return nil
}
