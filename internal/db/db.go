package db

import (
	"database/sql"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite3", "./favs.db")

	DB.SetMaxOpenConns(1)
	DB.SetMaxIdleConns(1)

	createTables()

	return DB, err
}

func createTables() {
	q := `CREATE TABLE IF NOT EXISTS favs (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		game_id TEXT NOT NULL,
		title TEXT NOT NULL
		);`

	DB.Exec(q)

	q = `CREATE TABLE IF NOT EXISTS stores (
		id INTEGER PRIMARY KEY,
		store_name TEXT NOT NULL,
		is_enabled BIT NOT NULL
	);`

	DB.Exec(q)
}
