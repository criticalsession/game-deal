package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/kyokomi/emoji/v2"
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

func UnixToDateString(unix int64) string {
	t := time.Unix(unix, 0)
	return t.Format("2006-01-02 15:04:05")
}

func PrintError(s string) {
	color.Red("%s %s\n\n", emoji.Sprintf(":red_exclamation_mark:"), s)
}

func PrintInfo(s string) {
	color.Yellow("! %s\n\n", s)
}

func GetIndexFromInput(s string) (int, error) {
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "]", "")

	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, errors.New("id should be a numerical value")
	}

	if id < 1 {
		return 0, errors.New("id should be greater than 0")
	}

	return id - 1, nil
}
