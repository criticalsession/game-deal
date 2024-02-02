package db

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/criticalsession/game-deal/internal/types/stores"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

type StoreRow struct {
	Id        int
	StoreName string
	IsEnabled bool
}

func GetStores() ([]StoreRow, error) {
	result := []StoreRow{}
	q := `SELECT * FROM stores;`
	rows, err := DB.Query(q)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		r := StoreRow{}
		err = rows.Scan(&r.Id, &r.StoreName, &r.IsEnabled)
		if err != nil {
			return result, err
		}

		result = append(result, r)
	}

	return result, nil
}

func GetEnabledStores() ([]StoreRow, error) {
	stores, err := GetStores()
	if err != nil {
		return nil, err
	}

	enabledStores := []StoreRow{}
	for _, s := range stores {
		if s.IsEnabled {
			enabledStores = append(enabledStores, s)
		}
	}

	return enabledStores, err
}

func GetStore(id int) (StoreRow, error) {
	q := `SELECT * FROM stores WHERE id = ?`
	res := StoreRow{}
	row := DB.QueryRow(q, id)
	err := row.Scan(&res.Id, &res.StoreName, &res.IsEnabled)

	if err == sql.ErrNoRows {
		return res, nil
	} else if err != nil {
		return res, err
	}

	return res, nil
}

func InsertStores(stores []stores.Store) error {
	for _, s := range stores {
		q := `INSERT INTO stores(id, store_name, is_enabled) VALUES(?, ?, 1)`
		_, err := DB.Exec(q, s.StoreID, s.StoreName)

		if err != nil {
			return err
		}
	}

	return nil
}

func SetStoreEnabled(id int, enabled bool) error {
	existing, err := GetStore(id)
	if err != nil {
		return err
	}

	if existing.Id != id {
		return errors.New("store does not exist in DB")
	}

	q := `UPDATE stores SET is_enabled = ? WHERE id = ?`
	_, err = DB.Exec(q, id, enabled)
	return err
}

func ClearStores() error {
	q := `TRUNCATE TABLE stores`
	_, err := DB.Exec(q)
	return err
}

func CheckStores(st []stores.Store) ([]stores.Store, error) {
	dbStores, err := GetStores()
	if err != nil {
		return []stores.Store{}, err
	}

	missingStores := []stores.Store{}
	for _, s := range st {
		found := false
		for _, e := range dbStores {
			id, err := strconv.Atoi(s.StoreID)
			if err != nil {
				return []stores.Store{}, err
			}

			if id == e.Id {
				found = true
				break
			}
		}

		if !found {
			missingStores = append(missingStores, s)
		}
	}

	return missingStores, nil
}
