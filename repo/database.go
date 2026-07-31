package repo

import (
	"context"
	"database/sql"
	"taskmanager/domain"
)

// Database is the standard database that is going to be used, made an interface for future implementations
type Database interface {
	SaveTask(name string, description string, status bool) (sql.Result, error)
	DeleteTask(name string) (sql.Result, error)
	Load(ctx context.Context) (map[string]domain.Task, error)
	Close() error
}
