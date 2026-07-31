package repo

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func createTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS tasks (
        title TEXT PRIMARY KEY,
        description TEXT,
        done INTEGER NOT NULL
    );`

	return db.Exec(sql)
}
