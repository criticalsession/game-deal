package utils

import (
	"fmt"
	"strconv"
)

func StringTo2fString(price string) (string, error) {
	f, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%.2f", f), nil
}
