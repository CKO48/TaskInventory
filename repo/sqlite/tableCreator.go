package repo

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

func createTable(db *sql.DB) (sql.Result, error) {
	sql := `CREATE TABLE IF NOT EXISTS tasks (
        name TEXT PRIMARY KEY,
        description TEXT,
        status INTEGER NOT NULL
    );`

	return db.Exec(sql)
}
