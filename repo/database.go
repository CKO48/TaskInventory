package repo

import "database/sql"

type Database interface {
	InsertTask(name string, description string) (sql.Result, error)
	DeleteTask(name string) (sql.Result, error)
	Close() error
}
