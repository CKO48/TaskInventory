package repo

import (
	"context"
	"database/sql"
	"taskmanager/domain"
	"taskmanager/repo"

	_ "modernc.org/sqlite"
)

type SQLiteRepo struct {
	db *sql.DB
}

/*
InitSQLiteDB creates or opens the database

Also ensures the table is created
*/
func InitSQLiteDB() (repo.Database, error) {
	db, err := sql.Open("sqlite", "./tasks.db")
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		title TEXT PRIMARY KEY,
		description TEXT,
		status BOOLEAN NOT NULL
	);`

	_, err = db.Exec(query)
	if err != nil {
		db.Close()
		return nil, err
	}

	return &SQLiteRepo{db: db}, nil
}

// SaveTask inserts a new task or updates its description/status if it already exists (UPSERT).
func (r *SQLiteRepo) SaveTask(name string, description string, status bool) (sql.Result, error) {
	query := `
	INSERT INTO tasks (title, description, status) 
	VALUES (?, ?, ?)
	ON CONFLICT(title) DO UPDATE SET
		description = excluded.description,
		status = excluded.status;`

	return r.db.Exec(query, name, description, status)
}

// DeleteTask deletes a task from the database based on its title.
func (r *SQLiteRepo) DeleteTask(name string) (sql.Result, error) {
	query := `DELETE FROM tasks WHERE title = ?;`
	return r.db.Exec(query, name)
}

// Load loads all tasks from the database and returns them as a map of title to Task.
func (r *SQLiteRepo) Load(ctx context.Context) (map[string]domain.Task, error) {
	query := `SELECT title, description, status FROM tasks;`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make(map[string]domain.Task)

	for rows.Next() {
		var title, description string
		var status bool

		if err := rows.Scan(&title, &description, &status); err != nil {
			return nil, err
		}

		// Se asume la función constructora domain.NewStandartTask(title, description, status)
		task := domain.NewStandartTask(title, description, status)
		tasks[title] = task
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Close closes the database connection.
func (r *SQLiteRepo) Close() error {
	return r.db.Close()
}
