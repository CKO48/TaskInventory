package repo

import (
	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

type SQLiteConnection struct {
	db *sql.DB
}

func InitSQLiteDB() (*SQLiteConnection, error) {
	db, err := sql.Open("sqlite", "./my.db")
	if err != nil {
		return nil, err
	}

	_, err = createTable(db)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &SQLiteConnection{db: db}, nil
}

func (conn *SQLiteConnection) InsertTask(name string, description string) (sql.Result, error) {
	sql := `INSERT INTO tasks (name, description, status) 
			VALUES (?, ?, ?)
			ON CONFLICT(name)
			DO UPDATE SET
				description = excluded.description,
				status = excluded.status;`
	return conn.db.Exec(sql, name, description, false)
}

func (conn *SQLiteConnection) DeleteTask(name string) (sql.Result, error) {
	sql := `DELETE FROM tasks WHERE name = ?;`
	return conn.db.Exec(sql, name)
}

func (conn *SQLiteConnection) Close() error {
	return conn.db.Close()
}
