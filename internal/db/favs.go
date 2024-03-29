package db

import (
	"errors"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type FavRow struct {
	Id     int
	GameId string
	Title  string
}

func AddFav(gameId, title string) error {
	existing, err := GetFavByGameId(gameId)
	if err == nil {
		return errors.New(existing.Title + " already in favorites")
	}

	q := `INSERT INTO favs (game_id, title) VALUES (?, ?);`

	_, err = DB.Exec(q, gameId, title)

	return err
}

func RemoveFav(id int) error {
	q := `DELETE FROM favs WHERE id = ?;`

	_, err := DB.Exec(q, id)

	return err
}

func GetFavs() ([]FavRow, error) {
	result := []FavRow{}

	q := `SELECT id, game_id, title FROM favs;`
	rows, err := DB.Query(q)

	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		r := FavRow{}

		err = rows.Scan(&r.Id, &r.GameId, &r.Title)
		if err != nil {
			return result, err
		}

		result = append(result, r)
	}

	return result, nil
}

func GetFavByIndex(id int) (FavRow, error) {
	results, err := GetFavs()
	if err != nil {
		return FavRow{}, err
	}

	if id >= len(results) {
		return FavRow{}, errors.New("id not found in fav list")
	}

	return results[id], nil
}

func GetFavByGameId(gameId string) (FavRow, error) {
	result := FavRow{}

	q := `SELECT id, game_id, title FROM favs WHERE game_id = ?;`
	row := DB.QueryRow(q, gameId)

	err := row.Scan(&result.Id, &result.GameId, &result.Title)
	if err != nil {
		return result, err
	}

	return result, nil
}
