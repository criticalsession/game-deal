package api

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func getData[T any](url string, client http.Client, result *T) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("bad status code")
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var unmarshaled T
	err = json.Unmarshal(dat, &unmarshaled)
	if err != nil {
		return err
	}

	*result = unmarshaled

	return nil
}
