package repo

import (
	"context"
	"database/sql"
	"taskmanager/domain"
)

type Database interface {
	SaveTask(name string, description string) (sql.Result, error)
	DeleteTask(name string) (sql.Result, error)
	Load(ctx context.Context) (map[string]domain.Task, error)
	Close() error
}
