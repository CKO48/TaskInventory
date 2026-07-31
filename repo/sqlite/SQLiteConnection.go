package repo

import (
	"database/sql"

	"context"
	"taskmanager/domain"

	_ "github.com/glebarez/go-sqlite"
)

type SQLiteConnection struct {
	db *sql.DB
}

// Initializes the SQLite database and creates the tasks table if it doesn't exist
func InitSQLiteDB() (*SQLiteConnection, error) {
	db, err := sql.Open("sqlite", "C:\\Users\\Kike\\Dev\\my.db")
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

// Load retrieves all tasks from the database and returns them as a map of task names to Task objects
func (conn *SQLiteConnection) Load(ctx context.Context) (map[string]domain.Task, error) {
	const query = `
		SELECT title, description, done
		FROM tasks;
	
	`

	rows, err := conn.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make(map[string]domain.Task)

	for rows.Next() {
		var title string
		var description string
		var done bool

		err := rows.Scan(
			&title,
			&description,
			&done,
		)
		if err != nil {
			return nil, err
		}

		tasks[title] = domain.NewStandartTask(title, description, done)
	}

	return tasks, nil
}

// adds a new task to the database or updates an existing task if it already exists
func (conn *SQLiteConnection) SaveTask(title string, description string, done bool) (sql.Result, error) {
	sql := `INSERT INTO tasks (title, description, done) 
			VALUES (?, ?, ?)
			ON CONFLICT(title)
			DO UPDATE SET
				description = excluded.description,
				done = excluded.done;`
	return conn.db.Exec(sql, title, description, done)
}

// Deletes a task from the database based on its title
func (conn *SQLiteConnection) DeleteTask(title string) (sql.Result, error) {
	sql := `DELETE FROM tasks WHERE title = ?;`
	return conn.db.Exec(sql, title)
}

// closes the database connection
func (conn *SQLiteConnection) Close() error {
	return conn.db.Close()
}
